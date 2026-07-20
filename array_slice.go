package main

import(
	"fmt"
)


func main() {
	// // -------------------------------ARRYA ---------------------------

	// var a [2]int
	// a[0]=1
	// a[1]=2
	// fmt.Println(a)


	// b:=[2]int{1,2}
	// fmt.Println(b)

	// c:=[...]int{1,2}
	// fmt.Println(c)

	// // 2D array
	// d:=[2][3]int{    // [r][c]
	// 	{1,2,3},     // 0
	// 	{4,5,6},     // 1
	//   // 0,1,2
	// }

	// fmt.Println(d[0][0])
	// fmt.Println(d[1][1])
	// fmt.Println(d[1][2])



	// -----------------Slices — dynamic, reference type ----------------------

	// s:=[]int{1,2,3,4,5}
	// // s:=make([]int,3,5)

	// s2:=append(s,6)
	// s2=append(s2,100)

	// fmt.Println(s)
	// fmt.Println(s2)



	// other:=[]int{9,0,8,8,8,8}
	// s3:=append(s,other...) //spread
	// fmt.Println(s3)

   // //deep copy
// 	dst:=make([]int,len(s))
// 	copy(dst,s)
// 	fmt.Println(dst)

// --------------------------------shallow copy

// a := []int{1, 2, 3}
// b := a

// b[0] = 100


// fmt.Println(a) // [100 2 3]
// fmt.Println(b) // [100 2 3]


// scires:=map[string]int{
// 	"hafiz":90,
// 	"nusayeb":100,
// }
// fmt.Println(scires["hafiz"])

// group:=map[string][]string{
// 	"admin":{"hafiz","nusayeb"}
// 	"user":{"Kazy","Noushin"}
// }

// config:= map[string]map[string]string{
// 	db:{ "host":"localhost","port":"27017"}
// }


a := []int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a) // [100 2 3]
fmt.Println(b) // [100 2 3]





}
