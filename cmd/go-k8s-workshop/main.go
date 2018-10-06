package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Hello, world")

	log.Fatal(http.ListenAndServe(":8888", nil))
}
