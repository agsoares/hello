package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!!!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running")
	http.ListenAndServe(":8080", nil)
}
