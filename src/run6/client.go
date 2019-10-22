package run6

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

func receive(i, topic int) (float64, error) {
	var start time.Time

	count := 0
	events := make(chan *sse.Event)
	topicStr := "chatroom_message" + strconv.Itoa(topic)
	url := "http://localhost:4000/_rig/v1/connection/sse?subscriptions=[{\"eventType\":\"" + topicStr + "\"}]"

	fmt.Println("Topic:", topicStr, "\t| Thread:", i)

	client := sse.NewClient(url)

	err := client.SubscribeChanRaw(events)

	if err != nil {
		return 0, err
	}

	for event := range events {
		if string(event.Event) != topicStr {
			continue
		}

		count++

		if count == 1 {
			start = time.Now()
		}

		if count == 1000 {
			elapsed := time.Since(start).Seconds()
			return elapsed, nil
		}

		if count%100 == 0 {
			go fmt.Println("Count:", count, "\t| Thread:", i, "\t| Topic:", topic)
		}

	}

	return 0, errors.New("")
}

func main() {
	var wg sync.WaitGroup

	goroutines, _ := strconv.Atoi(os.Getenv("CLIENTS"))
	topic := 1

	fmt.Println("Starting", goroutines, "goroutines")

	for i := 1; i <= goroutines; i++ {
		wg.Add(1)
		go func(i, topic int) {
			elapsed, err := receive(i, topic)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Thread", i, "finished in", elapsed, "s")

			wg.Done()
		}(i, topic)

		topic = topic + 1

		if topic > 100 {
			topic = 1
		}
	}

	fmt.Println("Waiting for goroutines to finish...")

	duration, err := time.ParseDuration(os.Getenv("TIMEOUT"))

	if err != nil {
		if loadtest.WaitTimeout(&wg, duration) {
			fmt.Println("Timed out waiting for wait group")
		} else {
			fmt.Println("Wait group finished")
		}
	} else {
		fmt.Println("Error while parsing timeout!")
	}
}
