// package main

// import (
//     "fmt"
// )

// /*

// 3. L — Liskov Substitution Principle (LSP):
// A subtype should be replaceable by its parent without breaking behavior.

// */

// func main() {
//     sparrow:=Sparrow{}
//     makeFly(sparrow) // passing child type Sparrow,argument expect parent type FlyingBird
// }
// // interface parent type
// type FlyingBird interface{
//     fly()
// }

// // concreat implementation - child type
// type Sparrow struct{}
// func (s Sparrow) fly(){
//     fmt.Println("Sparrow is flying")
// }

// func makeFly(fb FlyingBird){ // here Sparrow child type is replaced by parent type FlyingBird
//     fb.fly()
// }
