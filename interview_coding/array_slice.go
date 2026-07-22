// package main
// import(
// 	"fmt"
// )

// /*

// Interview takeaway:
// 	Arrays are copied. Slices are shared.
// 	Array assignment → creates a new copy.
// 	Slice creation ([:], [i:j]) → shares the same underlying array.

// 	A slice is just a view into an underlying array.
// 	capacity = len(arr) - start
// 	append() reuses the existing array if capacity allows.
// 	If capacity is exhausted, append() allocates a new array and copies the existing elements.
// 	That's why modifying one slice after append() can unexpectedly change the original array or other slices sharing the same backing array.


// 	| Operation                          | Shares underlying array?   | Change affects original? |
// | ---------------------------------- | -------------------------- | ------------------------ |
// | `b := a` (array)                   | ❌ No                       | ❌ No                     |
// | `s := arr[:]`                      | ✅ Yes                      | ✅ Yes                    |
// | `s2 := s[1:]`                      | ✅ Yes                      | ✅ Yes                    |
// | `append()` with available capacity | ✅ Yes                      | ✅ Yes                    |
// | `append()` when capacity is full   | ❌ No (new array)           | ❌ No                     |
// | Modify array                       | ✅ All slices see it        | ✅ Yes                    |
// | Modify slice                       | ✅ Underlying array changes | ✅ Yes                    |

// */

// func main(){
// 	// // array
// 	// arr1:=[3]int{1,2,3}
// 	// fmt.Println("arr1",arr1) //[1,2,3]

// 	// arr2:=arr1
// 	// arr2[0]=9

// 	// fmt.Println("arr2",arr2) //[9,2,3]
// 	// fmt.Println("arr1",arr1) //[1,2,3]

// 	// //slice
// 	// sl1:=arr1[0:1]
// 	// fmt.Println("sl1",sl1) //[1]

// 	// sl1[0]=100
// 	// fmt.Println("sl1",sl1) //[100]
// 	// fmt.Println("arr1",arr1) //[100,2,3]


// 	// // here append() reuses the existing array because capacity allows.
// 	// arr := [3]int{1, 2, 3}

// 	// s1 := arr[:2] //len=2, cap=3
// 	// s2 := append(s1, 99)

// 	// fmt.Println(arr) // [1,2,99]
// 	// fmt.Println(s1) // [1,2]
// 	// fmt.Println(s2) // [1,2,99]


//     // //  here append() allocates a new array and copies the existing elements, because capacity is exhauste, there is no room for new element
// 	// arr := [3]int{1, 2, 3}
// 	// s1 := arr[:3] // len=3, cap=3
// 	// s2 := append(s1, 99)

// 	// fmt.Println(arr) // [1,2,3]
// 	// fmt.Println(s1) // [1,2,3]
// 	// fmt.Println(s2) // [1,2,3,99]


// 	//// having extra room so nodify unerlying array
// 	// s := []int{1, 2, 3}
// 	// a := s[:2]
// 	// b := append(a, 100)

// 	// b[0] = 999

// 	// fmt.Println(s) // [999,2,100]
// 	// fmt.Println(a) // [999,2]
// 	// fmt.Println(b) // [999,2,100]


// 	// // not extra room foe new element so create new array and copy existing elements
// 	// s := []int{1, 2, 3}
// 	// a := s[:3]
// 	// b := append(a, 100)
// 	// b[0] = 999

// 	// fmt.Println(s) // [1,2,3]
// 	// fmt.Println(a) // [1,2,3]
// 	// fmt.Println(b) // [999,2,3,100]


// 	// ***************************** array element update through function ************************************
// 	/*

// 	| Type  | Modify element                                    | `append()`                                                                  |
// 	| ----- | ------------------------------------------------- | --------------------------------------------------------------------------- |
// 	| Array | ❌ Original not updated (unless passed as pointer) | N/A                                                                         |
// 	| Slice | ✅ Original updated                                | ❌ Caller's slice not updated unless you return the new slice or pass `*[]T` |


// 	*/
// 	// An array is passed by value in Go. So the function gets a copy of the array.
// 	// The original array is not updated.

// 	arr1 := [3]int{1, 2, 3}
// 	updateArr(arr1)
// 	fmt.Println(arr1) // [1,2,3]


// 	// If you want to modify the original array, pass a pointer
// 	arr2 := [3]int{1, 2, 3}
// 	updateArrPointer(&arr2)
// 	fmt.Println(arr2) // [9,2,3]


// 	// ****************************** slice element  update through function **********************************
// 	// If you modify an existing element
// 	// The original slice is updated.
// 	sl1 := []int{1, 2, 3}
// 	updateSlice(sl1)
// 	fmt.Println(sl1) // [9,2,3]


//     // The caller's slice is usually not updated, because append may create a new underlying array and only the local copy of the slice header changes.
// 	sl2 := []int{1, 2, 3}
// 	updateSliceWithAppend(sl2)
// 	fmt.Println(sl2) // [9,2,3]

// 	// Caller's slice not updated unless you return the new slice or pass *[]T
// 	sl3 := []int{1, 2, 3}
// 	sl3=updateSliceWithAppendAndReturn(sl3)
// 	fmt.Println("sl3",sl3) // [1,2,3,9]	
	
//     // pass *[]T
// 	sl4 := []int{1, 2, 3}
// 	updateSliceWithAppendAnPointer(&sl4)
// 	fmt.Println("sl4",sl4) // [1,2,3,9]




// }

// func updateArr(a [3]int){
//  a[0]=9
//  fmt.Println("inside updateArr function",a)
// }
// func updateArrPointer(a *[3]int){
//  a[0]=9
//   fmt.Println("inside updateArrPointer function",a)
// }

// func updateSlice(a []int){
//  a[0]=9
//    fmt.Println("inside updateSlice function",a)
// }

// func updateSliceWithAppend(a []int){
//  a=append(a,9)
//  fmt.Println("inside updateSliceWithAppend function",a)
// }

// func updateSliceWithAppendAndReturn(a []int) []int{
//  return append(a,9)
// }
// func updateSliceWithAppendAnPointer(a *[]int){
//  *a=append(*a,9)
//  fmt.Println("inside updateSliceWithAppendAnPointer function",a)
// }
