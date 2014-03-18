package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Getwd failed: %s", err))
	}

	fmt.Println("Starting server on", wd)
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir(wd))))
}
