// package main

// import (
//     "fmt"
// )

// /*

// 4. I — Interface Segregation Principle (ISP):
// Don't force clients to implement methods they don't need.

// */

// func main() {
// 	human:=Human{}
// 	human.work()
// 	human.eat()
// 	human.sleep()

// 	robot:=Robot{}
// 	robot.work()

    
// }
// // make interface with less number of method 
// type Worker interface{
//     work()
// }
// type Eater interface{
//     eat()
// }
// type Sleeper interface{
//     sleep()
// }

// // types
// type Human struct{}
// func(Human) work(){
// 	fmt.Println("Human can Work")
// }
// func(Human) eat(){
// 	fmt.Println("Human can Eat")
// }
// func(Human) sleep(){
// 	fmt.Println("Human can sleep")
// }

// type Robot struct{}
// func(Robot) work(){
// 	fmt.Println("Robot can work")
// }

