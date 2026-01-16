// channels: channels are a way for goroutines to communicate with each other and synchronize their execution. they can be thought of as pipes through which data can be sent and received.

package main

import (
	"fmt"
)

func main() {

	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)

	ch := make(chan int)

	// start a goroutine to send data into the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("sent:", i)
		}
		close(ch) // close the channel when done sending
	}()
}