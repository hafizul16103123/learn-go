package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two unbuffered channels for sending string messages.
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start a goroutine that waits for 2 seconds
	// and then sends a message to ch1.
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	// Start another goroutine that waits for 1 second
	// and then sends a message to ch2.
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	// Wait for the first channel that becomes ready.
	// Since ch2 sends after 1 second and ch1 after 2 seconds,
	// the message from ch2 will be received first.
	select {
	case msg := <-ch1:
		fmt.Println(msg)

	case msg := <-ch2:
		fmt.Println(msg)
	}
}