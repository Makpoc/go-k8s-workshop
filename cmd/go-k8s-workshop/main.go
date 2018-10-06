package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/makpoc/go-k8s-workshop/internal/diagnostics"
)

func main() {
	log.Print("Hello, world")

	router := mux.NewRouter()

	router.HandleFunc("/", hello)

	go func() {
		err := http.ListenAndServe(":8888", router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	diagnostics := diagnostics.NewDiagnostics()

	err := http.ListenAndServe(":9999", diagnostics)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
