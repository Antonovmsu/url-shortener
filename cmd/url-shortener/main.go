package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func saveURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form for saving URL"))
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	if r.PathValue("shortURL") == "google" {
		http.Redirect(w, r, "https://google.com", http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Custom 404: The resource you are looking for does not exist."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/save", saveURL)
	mux.HandleFunc("/{shortURL}", redirectURL)

	log.Print("Starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
