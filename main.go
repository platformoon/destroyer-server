package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	if err != nil {
		panic(err)
	}
}
