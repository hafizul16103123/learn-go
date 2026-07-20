// package main

// import(
// 	"fmt"
// )
// /*
// closure in Go is a function that captures and remembers variables from the surrounding scope,
// even after that outer function has returned.
// */

// func main(){
// myCounter:= counter()
// fmt.Println(myCounter())
// fmt.Println(myCounter())
// fmt.Println(myCounter())
// fmt.Println(myCounter())
// fmt.Println(myCounter())


// }
// func counter() func () int{
//   count:=0
//   return func() int{
//     count++
//     return count
//   }

// }