package main 
import (
"fmt"
)

/*
send goroutine to main go routine
*/
func main(){
	resultChan:= make(chan int,100)

	go sum(resultChan,2,9)

	sumRes:= <-resultChan
	fmt.Println("Sum Result: ", sumRes)

}

func sum(resultChan chan int,n int,n2 int){

	sum:=n+n2
	resultChan<- sum
}

