package models

import (
	"database/sql"
	"time"
)

type ShortURL struct {
	ID          int64
	OriginalURL string
	ShortCode   string
	Created     time.Time
	Expires     time.Time
}

type ShortURLModel struct {
	DB *sql.DB
}

func (m *ShortURLModel) Create(originalURL string, expires time.Time) (int, error) {
	return 0, nil
}

// TODO rethink what to return
func (m *ShortURLModel) Get(shortCode string) (*ShortURL, error) {
	return nil, nil
}
