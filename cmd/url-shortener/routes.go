package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /{shortURL}", app.redirectURL)
	mux.HandleFunc("GET /saveURL", app.saveURL)
	mux.HandleFunc("POST /saveURL", app.saveURLPost)

	return app.logRequest(commonHeaders(mux))
}
