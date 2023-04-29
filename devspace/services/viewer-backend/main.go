package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow World From viewer-backend!")
}

func main() {
	fmt.Println("Server Started on port: 9003")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":9003", nil)
	if err != nil {
		fmt.Println(err)
	}
}
