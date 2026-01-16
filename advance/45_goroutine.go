// goroutine is a lightweight thread managed by the Go runtime

package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello from goroutine:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sayHello() // start a new goroutine

	// main goroutine
	for i := 0; i < 5; i++ {
		fmt.Println("hello from main goroutine:", i)
		time.Sleep(150 * time.Millisecond)
	}
	// wait for a moment to let the goroutine finish
	time.Sleep(1 * time.Second)
}