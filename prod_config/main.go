package main

import (
	"fmt"
	"log"
	"config_module/config"
)


func main(){
	cfg,err:=config.MustLoad()
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(cfg.Env)
}
