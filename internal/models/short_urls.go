package models

import (
	"database/sql"
	"errors"
	"time"

	"url-shortener/internal/lib/random"
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

func (m *ShortURLModel) Insert(originalURL string, expires int) (ShortURL, error) {
	stmt := `INSERT INTO urls (original_url, short_code, created, expires)
	VALUES (?, ?, ?, ?)`

	u := ShortURL{
		OriginalURL: originalURL,
		ShortCode:   random.NewRandomString(6),
		Created:     time.Now(),
		Expires:     time.Now().Add(time.Duration(expires) * time.Hour * 24),
	}

	result, err := m.DB.Exec(stmt, u.OriginalURL, u.ShortCode, u.Created, u.Expires)
	if err != nil {
		return ShortURL{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return ShortURL{}, err
	}

	u.ID = id
	return u, nil
}

func (m *ShortURLModel) Get(shortCode string) (ShortURL, error) {
	stmt := `SELECT id, original_url, short_code, created, expires FROM urls 
	WHERE short_code = ?`

	var u ShortURL

	err := m.DB.QueryRow(stmt, shortCode).Scan(&u.ID, &u.OriginalURL, &u.ShortCode, &u.Created, &u.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ShortURL{}, ErrNoRecord
		}
		return ShortURL{}, err
	}

	return u, nil
}
