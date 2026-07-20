// package main

// import (
//     "fmt"
// )

/*
	1. S — Single Responsibility Principle (SRP):
	A type should have only one reason to change.

*/

// func main() {
//     user:=User{}

//     userRepo:= UserRepo{}
//     emailService:= EmailService{}
//     reportService:= ReportService{}

//     userRepo.SaveUser(&user,"Hafiz")
//     emailService.SendEmail(user.name)
//     reportService.GenerateReport(user.name)

// }

// type User struct{
//     name string
// }

// type UserRepo struct{}
// func(ur UserRepo) SaveUser(user *User,name string){
//     user.name=name
//     fmt.Println("Saved user",name)
// }

// type EmailService struct{}
// func (es EmailService) SendEmail(name string){
//     fmt.Println("Send Email to user",name)
// }

// type ReportService struct{}
// func(rs ReportService) GenerateReport(name string){
//     fmt.Println("Report Generated for user",name)
// }