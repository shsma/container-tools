package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow World From Publish-Backend!")
}

func main() {
	fmt.Println("Server Started on port: 9004")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":9004", nil)
	if err != nil {
		fmt.Println(err)
	}
}
