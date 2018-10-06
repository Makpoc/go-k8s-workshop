package diagnostics

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewDiagnostics() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/ready", ready)

	return r
}

func healthz(w http.ResponseWriter, r *http.Request) {
	log.Println("Got healthz request")
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
func ready(w http.ResponseWriter, r *http.Request) {
	log.Println("Got ready request")
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
