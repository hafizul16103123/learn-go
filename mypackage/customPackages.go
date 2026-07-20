package main
 import (
 	"fmt"
 )

func init(){

	config:=make(map[string]string)
	config["env"]="production"
	fmt.Println(config)
}

func SayHello() {
	println("Hello")
}