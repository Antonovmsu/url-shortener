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

	logger.Info("Starting server", slog.String("address", cfg.Address))

	err = http.ListenAndServe(cfg.Address, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
