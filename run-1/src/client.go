package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/r3labs/sse"
)

func main() {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	goroutines, _ := strconv.Atoi(os.Getenv("CLIENTS"))

	fmt.Println("Starting", goroutines, "goroutines")

	for i := 1; i <= goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			elapsed := receive(i)

			mutex.Lock()
			fmt.Println("Thread", i, "finished in", elapsed, "s")
			mutex.Unlock()

			wg.Done()
		}(i)
	}

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done")
}

func receive(i int) float64 {
	var start time.Time

	count := 1
	events := make(chan *sse.Event)
	client := sse.NewClient("http://localhost:4000/_rig/v1/connection/sse?subscriptions=[{\"eventType\":\"to_be_delivered\"}]")

	client.SubscribeChanRaw(events)
	duration, _ := time.ParseDuration("30s")

	go func(count *int) {
		for {
			time.Sleep(duration)

			if *count == 3 {
				return
			}

			fmt.Println(*count, i)
		}
	}(&count)

	for event := range events {
		if string(event.Event) != "to_be_delivered" {
			continue
		}

		count++

		if count == 2 {
			start = time.Now()
		}

		if count == 3 {
			elapsed := time.Since(start).Seconds()
			return elapsed
		}
	}

	return 0
}
