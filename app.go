package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gaspecia/go-to-do-list/middleware"
	"github.com/gorilla/mux"
)

type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := Info{
		Name:    "to-do-api",
		Version: "1.0.0",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.Use(
		middleware.LogginMiddleware(log.Default()),
	)

	router.HandleFunc("/", InfoHandler)

	httpServer := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        router,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}
