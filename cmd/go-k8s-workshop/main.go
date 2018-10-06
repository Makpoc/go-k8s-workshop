package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/makpoc/go-k8s-workshop/internal/diagnostics"
)

func main() {
	log.Print("Starting service...")

	blPort := os.Getenv("PORT")
	if blPort == "" {
		log.Fatal("PORT not provided")
	}
	diagPort := os.Getenv("DIAG_PORT")
	if diagPort == "" {
		log.Fatal("DIAG_PORT not provided")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	go func() {
		err := http.ListenAndServe(":"+blPort, router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	diagnostics := diagnostics.NewDiagnostics()
	err := http.ListenAndServe(":"+diagPort, diagnostics)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
