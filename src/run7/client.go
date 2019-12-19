package main

import (
	"fmt"
	"sync"

	loadtest ".."
)

func main() {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	goroutines, timeout := loadtest.GetEnv()
	fmt.Println("Starting", goroutines, "goroutines")

	for i := 1; i <= goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			elapsed, _ := loadtest.Receive(i, 100, 1, "deliver")

			mutex.Lock()
			fmt.Println("Thread", i, "finished in", elapsed, "s")
			mutex.Unlock()

			wg.Done()
		}(i)
	}

	fmt.Println("Waiting for goroutines to finish...")

	if loadtest.WaitTimeout(&wg, timeout) {
		fmt.Println("Timed out waiting for wait group")
	} else {
		fmt.Println("Wait group finished")
	}
}
