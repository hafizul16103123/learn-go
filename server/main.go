package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server on port 8888")
	})

	log.Println("Server listening on :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
