package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I'm hit.")
		fmt.Fprintf(w, "Atlas...")
	})

	go http.ListenAndServe(":8081", nil)
	http.ListenAndServe(":8082", nil)
}
