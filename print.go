package main

import (
    "fmt"
)

type User struct {
    Name string
    Age  int
}

func main() {
    fmt.Println("========= fmt.Print =========")
    fmt.Print("Hello ")
    fmt.Print("World")
    fmt.Println()

    fmt.Println("\n========= fmt.Println =========")
    fmt.Println("Hello", "Go", "Developers")

    fmt.Println("\n========= fmt.Printf =========")
    name := "Hafizul"
    age := 30
    fmt.Printf("Name: %s, Age: %d\n", name, age)
    fmt.Printf("Float: %.2f\n", 3.14159)
    fmt.Printf("Hex: %x\n", 255)
    fmt.Printf("Binary: %b\n", 10)

    fmt.Println("\n========= fmt.Sprintf =========")
    message := fmt.Sprintf("Welcome %s! Your age is %d.", name, age)
    fmt.Println(message)

    fmt.Println("\n========= fmt.Errorf =========")
    err := fmt.Errorf("user %d not found", 101)
    fmt.Println(err)

    fmt.Println("\n========= Format Verbs =========")

    num := 100
    pi := 3.14159
    active := true

    fmt.Printf("%%v  : %v\n", num)
    fmt.Printf("%%T  : %T\n", num)
    fmt.Printf("%%d  : %d\n", num)
    fmt.Printf("%%b  : %b\n", num)
    fmt.Printf("%%o  : %o\n", num)
    fmt.Printf("%%x  : %x\n", num)
    fmt.Printf("%%X  : %X\n", num)
    fmt.Printf("%%f  : %f\n", pi)
    fmt.Printf("%%.2f: %.2f\n", pi)
    fmt.Printf("%%e  : %e\n", pi)
    fmt.Printf("%%t  : %t\n", active)
    fmt.Printf("%%c  : %c\n", 65)
    fmt.Printf("%%q  : %q\n", "Hello")
    fmt.Printf("%%p  : %p\n", &num)

    fmt.Println("\n========= Struct Formatting =========")

    user := User{
        Name: "Alice",
        Age:  25,
    }

    fmt.Printf("%%v  : %v\n", user)
    fmt.Printf("%%+v : %+v\n", user)
    fmt.Printf("%%#v : %#v\n", user)

    fmt.Println("\n========= Sscanf =========")

    var id int
    var username string

    input := "101 Bob"

    fmt.Sscanf(input, "%d %s", &id, &username)

    fmt.Printf("ID=%d Name=%s\n", id, username)

    fmt.Println("\n========= Sprint Functions =========")

    a := fmt.Sprint("Go", " ", "Language")
    b := fmt.Sprintf("%d + %d = %d", 5, 3, 8)
    c := fmt.Sprintln("This", "contains", "newline")

    fmt.Println(a)
    fmt.Println(b)
    fmt.Print(c)

    fmt.Println("\n========= Complete =========")
}