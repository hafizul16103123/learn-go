package main

// ------------------------------------------------------ Without Go Routine ------------------------------------------------------

/**
  Without go routine it will take 20 seconds to complete the task because each worker will do the heavy work sequentially.
* doHeavyWork simulates a heavy task that a worker performs
*/

// import (
// 	"fmt"
// 	"time"
// )
// func doHeavyWork(id int) {
// 	fmt.Printf("Worker %d is doing heavy work...\n", id)
// 	// Simulate heavy work with a sleep
// 	time.Sleep(time.Second * 2)
// 	fmt.Printf("Worker %d has finished heavy work.\n", id)
// }
// func main() {
// 	start := time.Now()

// 	for i := 1; i <= 10; i++ {
// 		doHeavyWork(i)
// 	}

// 	fmt.Println("All workers have completed their tasks.")
// 	fmt.Printf("Total time taken: %v\n", time.Since(start))
// }

// ------------------------------------------------------ With Go Routine ------------------------------------------------------

/**
* With go routine it will take 2 seconds to complete the task because all workers will do the heavy work concurrently.
 */

import (
	"fmt"
	"sync"
	"time"
)
func doHeavyWork(id int,wg *sync.WaitGroup) {
  defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	fmt.Printf("Worker %d is doing heavy work...\n", id)
	// Simulate heavy work with a sleep
	time.Sleep(time.Second * 2)
	fmt.Printf("Worker %d has finished heavy work.\n", id)

}
func main() {
	start := time.Now()
	var wg sync.WaitGroup //initialize a WaitGroup to manage the synchronization of goroutines


	for i := 1; i <= 10; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go doHeavyWork(i,&wg)
	}


	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("All workers have completed their tasks.")
	fmt.Printf("Total time taken: %v\n", time.Since(start))
}