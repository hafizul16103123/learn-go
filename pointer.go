// package main

// import(
// 	"fmt"
// )
// /*

// */

// func main(){
//   	// x := 42
//     // p := &x          // p is *int — address of x
//     // fmt.Println(p)   // 0x... (memory address)
//     // fmt.Println(*p)  // 42 (dereference)

//     // *p = 100         // change x via pointer
//     // fmt.Println(x)   // 100



//     // // new() allocates zeroed memory
//     // q:=new(int)
//     // *q=500
//     // fmt.Println(q)
//     // fmt.Println(*q)



//     // // nil pointer
//     // var np *int
//     //  // fmt.Println(np)
//     // fmt.Println(*np)



// // Pointers with functions & structs

  
// // n:=5
// // dubVal(n)
// // fmt.Println(n) // 5

// // dubPointer(&n)
// // fmt.Println(n) // 10



// // with struct

// // c:=Counter{}
// // increment(c)
// // increment(c)
// // increment(c)
// // increment(c)
// // fmt.Println(c.val)


// cPointer:=&Counter{}
// incrementPointer(cPointer)
// incrementPointer(cPointer)
// incrementPointer(cPointer)
// incrementPointer(cPointer)
// fmt.Println(cPointer.val)



// }

//   func dubVal(n int){
//   	n *=2
//   }

//   func dubPointer(n *int){
//   	*n *=2
//   }

//   type Counter struct{
//   	val int
//   }

//   func incrementPointer(c *Counter){
//   	c.val++
//   }

//   func increment(c Counter){
//   	c.val++
//   }
