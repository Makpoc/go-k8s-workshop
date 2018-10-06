package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/makpoc/go-k8s-workshop/internal/diagnostics"
	"log"
	"net/http"
	"os"
	"time"
)

type serverConf struct {
	port    string
	handler http.Handler
	name    string
}

var calledN uint64

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
	router.HandleFunc("/count", metrics)

	diagRouter := diagnostics.NewDiagnostics()

	errChan := make(chan error, 2)

	serverConfs := []serverConf{
		{
			port:    blPort,
			handler: router,
			name:    "application",
		}, {
			port:    diagPort,
			handler: diagRouter,
			name:    "diagnostics",
		},
	}

	servers := make([]*http.Server, 2)

	for i, sc := range serverConfs {
		go func(conf serverConf, index int) {
			log.Printf("Starting %s server...", conf.name)
			srv := &http.Server{
				Addr:    ":" + conf.port,
				Handler: conf.handler,
			}
			servers[index] = srv
			err := srv.ListenAndServe()
			if err != nil {
				errChan <- err
			}
		}(sc, i)
	}

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Printf("Got an error: %v", err)

	case sig := <-interruptChan:
		log.Printf("Received %v signal. Stopping...", sig)
	}

	for _, srv := range servers {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := srv.Shutdown(ctx)
			if err != nil {
				log.Println(err)
			}
		}()
		log.Print("Server stopped")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&calledN, 1)
	log.Println("Got hello request")
	w.WriteHeader(200)
}

func metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"count": %d}`, calledN)))
}
