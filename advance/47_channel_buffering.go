// channel buffering: channel can be buffered. provide the buffer length as the second argument to make to initialize a buffered channel.

package main

import "fmt"

func main() {

    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}