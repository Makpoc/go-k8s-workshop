package diagnostics

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewDiagnostics() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/ready", ready)

	return r
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
func ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
