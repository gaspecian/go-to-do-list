package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gaspecian/go-to-do-list/config"
	"github.com/gaspecian/go-to-do-list/controller"
	"github.com/gaspecian/go-to-do-list/middleware"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(
		middleware.LogginMiddleware(log.Default()),
	)

	router.HandleFunc("/", controller.InfoHandler)

	httpServer := &http.Server{
		Addr:           config.LoadEnv().HttpServer.Addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        router,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}
