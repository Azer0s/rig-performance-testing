package loadtest

import (
	"sync"
	"time"
)

// WaitTimeout Wait for a wg to finish within a certain time, if the wg finishes before, false is returned
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
