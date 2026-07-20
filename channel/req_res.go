// package main
// import(
// 	"fmt"
// 	"sync"
// )

// func main(){
// 	var wg sync.WaitGroup
// 	jobChan:=make(chan int,2)
// 	resChan:=make(chan string,2)

// 	wg.Add(2)
// 	go worker(1,jobChan,resChan,&wg)
// 	go worker(2,jobChan,resChan,&wg)

//     // send data to job channel
// 	jobChan<- 100
// 	jobChan<- 200
// 	close(jobChan)


// 	// receive data from result Channel

// 	fmt.Println(<-resChan)
// 	fmt.Println(<-resChan)
	
// 	/*
// 	This is best because it lets the main goroutine keep consuming results 
// 	while a separate goroutine safely waits for all workers to finish and 
// 	closes the channel exactly once, preventing both deadlocks and “send on closed channel” panics.
// 	*/
// 	go func(){	
// 		wg.Wait()
// 		close(resChan)
// 	}()
// }

// func worker(id int,jobChan <-chan int,resChan chan<- string, wg *sync.WaitGroup){
// 	defer wg.Done()

// 	// receive data from job channel
// 	job:= <-jobChan

// 	result:= fmt.Sprintf(
// 		"Worker %d process %d",id,job,
// 	)

// 	// send data to result chan
// 	resChan<- result

// }

