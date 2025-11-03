package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, User!"))
}

func saveURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form for saving URL"))
}

func saveURLPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save new URL"))
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	// TODO: Add URL validation
	if r.PathValue("shortURL") == "google" {
		http.Redirect(w, r, "https://google.com", http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Custom 404: The resource you are looking for does not exist."))
}

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
