package main

import (
	"log/slog"
	"net/http"
	"os"

	"url-shortener/internal/config"
)

func main() {
	// TODO: Specify log output
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// TODO: Read config path from env
	cfg, err := config.MustLoad("./configs/local.yaml")
	if err != nil {
		logger.Error(err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /{shortURL}", redirectURL)
	mux.HandleFunc("GET /saveURL", saveURL)
	mux.HandleFunc("POST /saveURL", saveURLPost)

	logger.Info("Starting server", slog.String("address", cfg.Address))

	err = http.ListenAndServe(cfg.Address, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
