package main_godotenv

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/",hello)
	http.ListenAndServe(":8000",nil)
}
func hello(w http.ResponseWriter,r *http.Request){
	fmt.Println("Method:", r.Method)

	fmt.Println("Path:", r.URL.Path)

	env:=os.Getenv("ENV")
	fmt.Fprintln(w,"Hello Go",env,"OK")
}