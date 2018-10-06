package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Print("Hello, world")

	router := mux.NewRouter()

	router.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(":8888", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
