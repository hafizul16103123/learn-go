// package main

// import(
// 	"fmt"
// 	"encoding/json"
// )
// /*
// Struct embedding in Go promotes inner struct fields to the outer struct,
//  but if multiple embedded structs have the same field name, 
//  you must explicitly specify which struct you mean. otherwise AMBIGUOUS will occur.
// */

// type Address struct {
//     Name    string
//     City    string
//     Country string
// }

// type Person struct{
//  Name string `json:"name"`
//  Age int     `json:"age"`
// }

// type Employee struct {
//     Person          // embedded — fields promoted
//     Address         // embedded
//     Company string
//     Salary  float64
// }

// func main() {
//     e := Employee{
//         Person:  Person{Name: "Alice", Age: 30},
//         Address: Address{Name:"NYC",City: "NYC", Country: "US"},
//         Company: "Acme",
//         Salary:  75000,
//     }

//     // Promoted fields
//     fmt.Println(e.Person.Name)     // from Person
//     fmt.Println(e.City)     // from Address
//     fmt.Println(e.Company)



//     p2:=Person{Name: "Hafiz", Age: 30}

//     data,_:=json.Marshal(p2)
//     fmt.Println(string(data))


// }