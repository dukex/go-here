package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	here, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("failed: %s", err))
	}

	fmt.Println("Starting server on", here)
	fmt.Println("Open http://localhost:8080")
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir(here))))
}
