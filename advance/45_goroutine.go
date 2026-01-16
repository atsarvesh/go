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

func f(from string) {
	for i := range 3{
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct call") // direct call
	go f("goroutine call") // call as a goroutine

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	go sayHello() // start a new goroutine

	// main goroutine
	for i := 0; i < 5; i++ {
		fmt.Println("hello from main goroutine:", i)
		time.Sleep(150 * time.Millisecond)
	}
	// wait for a moment to let the goroutine finish
	time.Sleep(1 * time.Second)
}