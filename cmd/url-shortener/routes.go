package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /{shortURL}", app.redirectURL)
	mux.HandleFunc("GET /saveURL", app.saveURL)
	mux.HandleFunc("POST /saveURL", app.saveURLPost)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
