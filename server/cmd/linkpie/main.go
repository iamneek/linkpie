package main

import (
	"log"
	"net/http"

	"github.com/iamneek/linkpie/internal/router"
)

func main() {
	mux := router.New()

	server := &http.Server{
		Addr:    ":5000",
		Handler: mux,
	}

	log.Println("Starting server on :5000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
