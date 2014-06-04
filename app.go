package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var port = flag.Int("p", 8080, "-p 3000")

func main() {
	flag.Parse()

	portString := strconv.Itoa(*port)

	here, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("failed: %s", err))
	}

	fmt.Println("Starting server on", here)
	fmt.Println("Open http://localhost:" + portString)
	panic(http.ListenAndServe(":"+portString, http.FileServer(http.Dir(here))))
}
