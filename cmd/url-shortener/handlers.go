package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"url-shortener/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, User!"))
}

func (app *application) saveURL(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/html/pages/create.html")
}

func (app *application) saveURLPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	original_url := r.PostForm.Get("url")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	fieldErrors := make(map[string]string)

	if _, err := url.ParseRequestURI(original_url); err != nil {
		fieldErrors["original_url"] = "URL is not valid"
	}

	if expires < 1 {
		fieldErrors["expires"] = "Expires must be at least 1 day"
	}

	if len(fieldErrors) > 0 {
		fmt.Fprint(w, fieldErrors)
		return
	}

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
