package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/r3labs/sse"
)

func receive() (float64, error) {
	var start time.Time

	count := 1
	events := make(chan *sse.Event)
	client := sse.NewClient("http://localhost:4000/_rig/v1/connection/sse?subscriptions=[{\"eventType\":\"to_be_delivered\"}]")

	client.SubscribeChanRaw(events)

	for event := range events {
		if string(event.Event) != "to_be_delivered" {
			continue
		}

		count++

		if count == 2 {
			start = time.Now()
			fmt.Println("Starting timer")
		}

		if count == 3 {
			elapsed := time.Since(start).Seconds()
			return elapsed, nil
		}
	}

	return 0, errors.New("")
}

func main() {
	elapsed, _ := receive()
	fmt.Println(elapsed, "seconds")
}
