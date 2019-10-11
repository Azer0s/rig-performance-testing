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

func receive(i, topic int) (float64, error) {
	var start time.Time

	count := 1
	events := make(chan *sse.Event)
	topicStr := "chatroom_message" + strconv.Itoa(topic)
	url := "http://localhost:4000/_rig/v1/connection/sse?subscriptions=[{\"eventType\":\"" + topicStr + "\"}]"

	fmt.Println(url)

	client := sse.NewClient(url)

	client.SubscribeChanRaw(events)

	for event := range events {
		if string(event.Event) != topicStr {
			continue
		}

		count++

		if count == 2 {
			start = time.Now()
		}

		if count%100 == 0 {
			go fmt.Println("Count:", count, "| Thread:", i, "| Topic:", topic)
		}

		if count == 1000 {
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
	topic := 1

	fmt.Println("Starting", goroutines, "goroutines")

	for i := 1; i <= goroutines; i++ {
		wg.Add(1)
		go func(i, topic int) {
			elapsed, _ := receive(i, topic)

			mutex.Lock()
			fmt.Println("Thread", i, "finished in", elapsed, "s")
			mutex.Unlock()

			wg.Done()
		}(i, topic)

		topic = topic + 1

		if topic > 100 {
			topic = 1
		}
	}

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done")
}
