package main

import (
	"log"
	"net/http"

	"url-shortener/internal/config"
)

func main() {
	cfg, err := config.MustLoad("./configs/local.yaml")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /{shortURL}", redirectURL)
	mux.HandleFunc("GET /saveURL", saveURL)
	mux.HandleFunc("POST /saveURL", saveURLPost)

	log.Printf("Starting server on %s", cfg.Address)

	err = http.ListenAndServe(cfg.Address, mux)
	log.Fatal(err)
}
