// package main

// import(
// 	"fmt"
// )

// func main() {
// 	/*
// 	Shallow Copy:
// 	A shallow copy copies only the top-level value.
// 	If the value contains pointers (like slice/map), both copies share the same underlying data.
// 	🔹 Shallow copy:
// 	Copies structure only
// 	Shares underlying data
// 	Faster but unsafe for shared mutation

// 	👉 In Go, assignment creates a shallow copy; slices, maps, and pointers share underlying data, 
// 	so deep copy must be implemented manually using copy() or explicit struct reconstruction.
	
// 	🔹When Shallow Copy Happens:
// 	Slice/Map/Pointer assignment
// 	Struct assignment (if struct contains slices/maps/pointers)
	
// 	🔹 When Deep Copy Happens:
// 	Arrays (default behavior)
// 	Slice deep copy using copy()
// 	Manual struct deep copy
// 	*/



// // // assign share same underlying array
// // a := []int{1, 2, 3}
// // b := a

// // b[0] = 100

// // fmt.Println(a) // [100 2 3]
// // fmt.Println(b) // [100 2 3]




// // // Slice inside struct is shared
// // type User struct {
// // 	name string
// // 	data []int
// // }

// // u1 := User{"A", []int{10, 20}}
// // fmt.Println(u1)
// // u2 := u1

// // u2.data[0] = 999
// // u2.name = "B"

// // fmt.Println(u1) // [999 20]
// // fmt.Println(u2) // [999 20]

// /*
// Deep Copy

// A deep copy duplicates everything, including underlying data.
// 🔹 Deep copy:
// 	Copies everything
// 	Independent memory
// 	Safer for concurrency and isolation

// */

// // //✔ Independent memory
// // a := []int{1, 2, 3}

// // b := make([]int, len(a))
// // copy(b, a)

// // b[0] = 100

// // fmt.Println(a) // [1 2 3]
// // fmt.Println(b) // [100 2 3]



// // //Struct deep copy
// // type User struct {
// // 	name string
// // 	data []int
// // }

// // u1 := User{"A", []int{10, 20}}
// // fmt.Println(u1)
// // u2 := User{
// // 	name:u1.name,
// // 	data:append([]int{},u1.data...),
// // }

// // u2.data[0] = 999
// // u2.name = "B"

// fmt.Println(u1) // [999 20]
// fmt.Println(u2) // [999 20]




// }
