// package main

// import(
// 	"fmt"
// )


// func main() {

// 	// nums:=[]int{1,5,8}

// 	// addition:=add(5,8)
// 	// fmt.Println(addition)

// 	// divRes,err:=div(8,0)
// 	// fmt.Println(divRes,err)

// 	// min,max:=minMax([]int{1,5,8})
// 	// fmt.Println(min,max)

// 	// sumRes :=sum(nums...) //spread slice
// 	// fmt.Println(sumRes)

// 	mulFuncDef:= func(a,b int) int{
// 		return a*b
// 	}

// 	res1:=apply(mulFuncDef,5,5)
// 	fmt.Println(res1)

// }
// // single return value
// func add(a,b int) int{
// 	return a+b
// }
// // Multiple return values
// func div(a,b int) (int,error){
// 	if(b==0){return 0,fmt.Errorf("division by 0")}
// 	return a/b,nil
// }
// //// Named return values
// func minMax(nums []int)(min,max int){
// 	min,max=nums[0],nums[0]

// 	for _,v:=range nums[1:]{
// 		if v<min {min=v}
// 		if v>max {max=v}
// 	}
// 	return
// }

// //Variadic functions
// func sum(nums ...int) int{
// 	total:=0
// 	for _,v:= range nums{
// 		total+=v
// 	}
// 	return total
// }
// /*
// f func(int,int)int // function type/signature
// func (a,b int)int{return a+b} // function defination
// */

// //Functions as values
// func apply(f func(int,int)int,a,b int) int{
// 	return f(a,b)
// }