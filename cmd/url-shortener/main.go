package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"url-shortener/internal/config"
	"url-shortener/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger     *slog.Logger
	short_urls *models.ShortURLModel
}

func main() {
	// TODO: Use environment variables
	config_path := flag.String("cfg", "./configs/local.yaml", "Path to config file")
	dsn := flag.String("dsn", "web:pass@/urlshortener?parseTime=True", "MySQL DSN")
	flag.Parse()

	// TODO: Specify log output
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger:     logger,
		short_urls: &models.ShortURLModel{DB: db},
	}

	cfg, err := config.MustLoad(*config_path)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Starting server", slog.String("address", cfg.Address))

	err = http.ListenAndServe(cfg.Address, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
