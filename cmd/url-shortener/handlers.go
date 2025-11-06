package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, User!"))
}

func (app *application) saveURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form for saving URL"))
}

func (app *application) saveURLPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save new URL"))
}

func (app *application) redirectURL(w http.ResponseWriter, r *http.Request) {
	// TODO: Add URL validation
	if r.PathValue("shortURL") == "google" {
		http.Redirect(w, r, "https://google.com", http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Custom 404: The resource you are looking for does not exist."))
}
