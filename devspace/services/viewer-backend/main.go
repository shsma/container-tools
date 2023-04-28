package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow World From DevSpace!")
}

func main() {
	fmt.Println("Server Started on port: 9003")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":9003", nil)
}
