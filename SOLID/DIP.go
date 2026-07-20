// package main

// import (
//     "fmt"
// )

// /*

// 5. D — Dependency Inversion Principle (DIP):
// Depend on abstractions, not concrete implementations.

// */

// func main() {
	
// 	userService:=UserService{
// 		db:MySQL{},
// 	}
// 	userService.registerUser()

// 	userService2:=UserService{
// 		db:Mongo{},
// 	}
// 	userService2.registerUser()
    
// }

// type Database interface{
// 	save()
// }


// type UserService struct{
// 	db Database
// }

// func(us UserService) registerUser(){
// 	us.db.save()
// }

// type MySQL struct{}
// func(MySQL) save(){
// 	fmt.Println("Saved to MySQL")
// }

// type Mongo struct{}
// func(Mongo) save(){
// 	fmt.Println("Saved to Mongo")
// }

