package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/r3labs/sse"
)

func receive(i int) (float64, error) {
	var start time.Time

	count := 1
	events := make(chan *sse.Event)
	client := sse.NewClient("http://localhost:4000/_rig/v1/connection/sse?subscriptions=[{\"eventType\":\"chatroom_message\"}]")

	client.SubscribeChanRaw(events)

	for event := range events {
		if string(event.Event) != "chatroom_message" {
			continue
		}

		count++

		if count == 2 {
			start = time.Now()
		}

		if count%1000 == 0 {
			go fmt.Println(count, i)
		}

		if count == 100000 {
			elapsed := time.Since(start).Seconds()
			return elapsed, nil
		}
	}

	return 0, errors.New("")
}

func main() {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	goroutines, _ := strconv.Atoi(os.Getenv("CLIENTS"))

	fmt.Println("Starting", goroutines, "goroutines")

	for i := 1; i <= goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			elapsed, _ := receive(i)

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