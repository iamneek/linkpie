package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home"))
	})

	fmt.Println("Starting server: http://localhost:5000")
	if err := http.ListenAndServe(":5000", mux); err != nil {
		fmt.Println("Error starting server")
	}
}
