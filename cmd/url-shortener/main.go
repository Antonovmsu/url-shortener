package main

import (
	"log/slog"
	"net/http"
	"os"

	"url-shortener/internal/config"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// TODO: Specify log output
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	// TODO: Read config path from env
	cfg, err := config.MustLoad("./configs/local.yaml")
	if err != nil {
		logger.Error(err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /{shortURL}", app.redirectURL)
	mux.HandleFunc("GET /saveURL", app.saveURL)
	mux.HandleFunc("POST /saveURL", app.saveURLPost)

	logger.Info("Starting server", slog.String("address", cfg.Address))

	err = http.ListenAndServe(cfg.Address, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
