package main

import (
	"encoding/json"
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

	vals := make(map[int]bool)

	for i := 0; i < 10000; i++ {
		vals[i] = false
	}

	for event := range events {
		if string(event.Event) != "deliver" {
			continue
		}

		var f interface{}
		json.Unmarshal(event.Data, &f)

		if f == nil {
			f = map[string]interface{}{
				"data": "NOTHING",
			}
		}

		m := f.(map[string]interface{})
		mData := m["data"].(string)

		val, err := strconv.Atoi(mData)

		if err != nil {
			if mData == "DONE" {
				elapsed := time.Since(start).Seconds()

				for k, v := range vals {
					if v == false {
						fmt.Println("Unreceived", k, "by", i)
					}
				}

				return elapsed, nil
			}
		} else {
			vals[val] = true
		}

		count++

		if count == 2 {
			start = time.Now()
		}
	}

	return 0, errors.New("")
}

func main() {
	var wg sync.WaitGroup

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

	for i := 1; i <= goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			receive(i)
			wg.Done()
		}(i)
	}

	fmt.Println("Connected!")

	if loadtest.WaitTimeout(&wg, os.Getenv("TIMEOUT")) {
		fmt.Println("Timed out waiting for wait group")
	} else {
		fmt.Println("Wait group finished")
	}
}
