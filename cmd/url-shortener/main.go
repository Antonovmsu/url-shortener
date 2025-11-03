package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /{shortURL}", redirectURL)
	mux.HandleFunc("GET /saveURL", saveURL)
	mux.HandleFunc("POST /saveURL", saveURLPost)

	log.Print("Starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
