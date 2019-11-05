package loadtest

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/r3labs/sse"
)

// Receive & time measurement loop
func Receive(i, amount, printSteps int, sub string) (float64, error) {
	var start time.Time

	count := 1
	events := make(chan *sse.Event)
	client := sse.NewClient("http://" + os.Getenv("RIG_HOST") + ":4000/_rig/v1/connection/sse?subscriptions=[{\"eventType\":\"" + sub + "\"}]")

	client.SubscribeChanRaw(events)

	for event := range events {
		if string(event.Event) != "deliver" {
			continue
		}

		count++

		if count == 2 {
			start = time.Now()
		}

		if count%printSteps == 0 {
			go fmt.Println("Count:", count, "\t| Thread:", i, "\t| Topic:", sub)
		}

		if count == amount {
			elapsed := time.Since(start).Seconds()
			return elapsed, nil
		}
	}

	return 0, errors.New("")
}

// GetEnv Checks env and returns amount of goroutines to start
func GetEnv() (int, string) {
	if os.Getenv("RIG_HOST") == "" {
		fmt.Println("RIG_HOST environment variable required!")
		os.Exit(1)
	}

	wait, err := time.ParseDuration(os.Getenv("WAIT"))

	if err != nil {
		fmt.Println("WAIT environment variable required!")
		os.Exit(1)
	}

	if os.Getenv("CLIENTS") == "" {
		fmt.Println("CLIENTS environment variable required!")
		os.Exit(1)
	}

	goroutines, err := strconv.Atoi(os.Getenv("CLIENTS"))

	if err != nil {
		fmt.Println("Error while parsing CLIENTS!")
		os.Exit(1)
	}

	if os.Getenv("TIMEOUT") == "" {
		fmt.Println("TIMEOUT environment variable required!")
		os.Exit(1)
	}

	time.Sleep(wait)
	return goroutines, os.Getenv("TIMEOUT")
}
