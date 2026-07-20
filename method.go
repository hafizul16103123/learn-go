// package main

// import(
// 	"fmt"
//   "math"
// )
// /*

// */

// func main(){

//   c:=Circle{radius:5}
//   area:=c.Area()
//   fmt.Println(area)

//   c.Scale()
//   fmt.Println(c.radius)

// }
// type Circle struct{
//   radius float64
// }
// // Value receiver — read-only, works on copy
// func (c Circle) Area() float64{
//   return math.Pi * c.radius * c.radius
// }

// // Pointer receiver — can modify struct
// func (c *Circle) Scale(){
//   c.radius *= 2
// }



