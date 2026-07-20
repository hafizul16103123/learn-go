package main 
import (
"fmt"
"math/rand"
)

/*
send data to main goroutine to child go routine
*/
func main(){
	numChan :=make(chan int,100)

	go numConsumer(numChan)

	for{
		numChan<-rand.Intn(100)
	}
}

func numConsumer(numChan <-chan int){
	for ch:=range numChan{
		fmt.Println("Process number: ",ch)
	}

}