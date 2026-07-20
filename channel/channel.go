// package main

// import "fmt"
// /*
// Channels are FIFO (First In, First Out), like a queue.
// Unbuffered Channel:
// 	ch := make(chan int)
// 	Behavior:
//  	No storage/buffer.
// 	Sender and receiver must be ready at the same time.
// 	Send blocks until a receiver receives.
// 	Receive blocks until a sender sends.
// 	The send and receive synchronize with each other.
// 	Use when you need synchronization.

// 	To Prevent deadlock of a unbuffered channel
// 		- convert it to buffered channel 
// 		- use goroutine

// Buffered Channel:
// 	ch := make(chan int, 3)
// 	Behavior:
// 	Has internal storage.
// 	Sender can send until the buffer is full.
// 	Receiver can receive later.
// 	Use when producers and consumers run at different speeds.
// 	Workers can process jobs later while producers keep adding jobs.

// 	Common use cases:
// 	Worker pools
// 	Job queues
// 	Event processing
// 	Background tasks




// | Operation       | Unbuffered                  | Buffered                      |
// | --------------- | --------------------------- | ----------------------------- |
// | Send            | Blocks until receiver ready | Blocks only when buffer full  |
// | Receive         | Blocks until sender ready   | Blocks only when buffer empty |
// | Storage         | No                          | Yes                           |
// | Synchronization | Strong                      | Loose                         |

// */
// func main() {
//  ch:=make(chan int)

//  // send data to channel
//  //ch<- 100 // this will block the execution

//  go func(){
//  	ch<- 100
//  }()


//  // Receive data from channel
//  number:=<-ch

//  fmt.Println(number)

// }
