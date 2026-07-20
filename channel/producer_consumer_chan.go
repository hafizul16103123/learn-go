// package main
// import(
// 	"fmt"
// 	"sync"
// )

// /*
// Sending data from main goroutine to child goroutine
// Correct Pattern (one channel, two roles)
// */
// func main(){
// 	var wg sync.WaitGroup
// 	// use same channel as producer and consumer
// 	ch:= make(chan int,100)

// 	wg.Add(2)
// 	go producer(ch,&wg)
// 	go consumer(ch,&wg)

// 	wg.Wait()
// }

// func producer(producerCh chan<- int,wg *sync.WaitGroup){
// 	defer wg.Done()

// 	producerCh<-50
// 	producerCh<-100
// 	close(producerCh) // ✅ IMPORTANT
// }

// func consumer(consumerCh <-chan int, wg *sync.WaitGroup){
// 	defer wg.Done()

// 	for cCh:=range consumerCh{
// 		fmt.Println("Procerss",cCh)
// 	}
// }

