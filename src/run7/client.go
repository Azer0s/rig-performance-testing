package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	loadtest ".."

	"github.com/r3labs/sse"
)

func receive(i int) (float64, error) {
	var start time.Time

	count := 1
	events := make(chan *sse.Event)
	client := sse.NewClient("http://" + os.Getenv("RIG_HOST") + ":4000/_rig/v1/connection/sse?subscriptions=[{\"eventType\":\"deliver\"}]")

	client.SubscribeChanRaw(events)

	for event := range events {
		if string(event.Event) != "deliver" {
			continue
		}

		count++

		if count == 2 {
			start = time.Now()
		}

		go fmt.Println(count, i)

		if count == 100 {
			elapsed := time.Since(start).Seconds()
			return elapsed, nil
		}
	}

	return 0, errors.New("")
}

func main() {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	if os.Getenv("RIG_HOST") == "" {
		fmt.Println("RIG_HOST environment variable required!")
		os.Exit(1)
	}

	wait, err := time.ParseDuration(os.Getenv("WAIT"))

	if err != nil {
		fmt.Println("WAIT environment variable required!")
		os.Exit(1)
	}

	time.Sleep(wait)

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

	if loadtest.WaitTimeout(&wg, os.Getenv("TIMEOUT")) {
		fmt.Println("Timed out waiting for wait group")
	} else {
		fmt.Println("Wait group finished")
	}
}
