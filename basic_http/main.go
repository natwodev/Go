package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to Golang Web Server!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
