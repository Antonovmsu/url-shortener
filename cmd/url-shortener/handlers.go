package main

import (
	"errors"
	"fmt"
	"net/http"

	"url-shortener/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, User!"))
}

func (app *application) saveURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form for saving URL"))
}

func (app *application) saveURLPost(w http.ResponseWriter, r *http.Request) {
	original_url := "https://ya.ru"
	expires := 1

	short_url, err := app.short_urls.Insert(original_url, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Info(
		"new URL created",
		"id", short_url.ID,
		"originalUrl", short_url.OriginalURL,
		"shortCode", short_url.ShortCode,
		"expires", short_url.Expires)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "URL created: http://%s/%s", r.Host, short_url.ShortCode)
}

func (app *application) redirectURL(w http.ResponseWriter, r *http.Request) {
	resURL, err := app.short_urls.Get(r.PathValue("shortURL"))
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	http.Redirect(w, r, resURL.OriginalURL, http.StatusTemporaryRedirect)
}
