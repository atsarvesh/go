/*
	channel synchronization:
	means that both sender and receiver should be ready to send and receive data
*/

package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
}