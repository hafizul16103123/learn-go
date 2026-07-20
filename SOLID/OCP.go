// package main

// import (
//     "fmt"
// )

// /*
// 2. O — Open/Closed Principle (OCP):
//     Open for extension, closed for modification.
//     You should be able to add new functionality without changing existing code.
// */

// func main() {
//     paymentProcessor:=PaymentProcessor{}

//     card:=Card{}
//     bKash:=Bkash{}
//     rocket:=Rocket{}

//     paymentProcessor.checkout(card)
//     paymentProcessor.checkout(bKash)
//     paymentProcessor.checkout(rocket)

// }

// // interface
// type PaymentMethod interface{
//     pay()
// }
// // Payment Processor
// type PaymentProcessor struct{}
// func(pp PaymentProcessor) checkout(pm PaymentMethod){
//     pm.pay()
// }


// // Concreat implementaion of PaymentMethod interface
// type Card struct{}
// func(c Card) pay(){
//     fmt.Println("Payment by CARD")
// }

// type Bkash struct{}
// func(b Bkash) pay(){
//     fmt.Println("Payment by Bkash")
// }

// type Rocket struct{}
// func(r Rocket) pay(){
//     fmt.Println("Payment by Rocket")
// }


