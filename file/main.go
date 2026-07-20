// package main

// import (
//     "os"
//     // "io"
//     "fmt"
// )

// /*

// */

// func main() {
// 	// file,err:=os.Create("help.txt") //Create a file
// 	// defer file.Close() //Close the file

// 	// if(err !=nil) {panic("file not created")}

// 	// file.Write([]byte("Hello")) //Write to a file

// 	fileStat,fileExistsErr:=os.Stat("helpv2.txt") // check file exists or not
// 	if os.IsNotExist(fileExistsErr){
// 		fmt.Println("help.txt does not exist")
// 	}
// 		fmt.Println(fileStat)
	

// 	// err:=os.Rename("help.txt","helpv2.txt")
// 	// 	if err != nil {
// 	// 	fmt.Println("Rename failed:", err)
// 	// 	return
// 	// }

// 	// data, err := io.ReadAll(file)
// 	// fmt.Println(data)
	
    
// }
