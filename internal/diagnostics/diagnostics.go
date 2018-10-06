package diagnostics

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/mux"
)

func NewDiagnostics() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/ready", ready)

	return r
}
var calledReady uint64

func healthz(w http.ResponseWriter, r *http.Request) {
	log.Println("Got healthz request")
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
func ready(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&calledReady, 1)
	log.Println("Got ready request")
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
