package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hellow From Devspace publish-backend!")
	if err != nil {
		panic("Error while writing response")
	}
}

func main() {
	fmt.Println("Server Started on port: 9004")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":9004", nil)
	if err != nil {
		fmt.Println(err)
	}
}
