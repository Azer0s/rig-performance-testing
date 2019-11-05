package main

import (
	"fmt"

	loadtest ".."
)

func main() {
	elapsed, _ := loadtest.Receive(0, 3, 1, "to_be_delivered")
	fmt.Println(elapsed, "seconds")
}
