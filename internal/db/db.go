package db

import "database/sql"

// CreateTable creates the urls table in the database if it doesn't exist.
func CreateTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS urls (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            short_url TEXT UNIQUE NOT NULL,
            original_url TEXT NOT NULL
		);`
    _, err := db.Exec(query)
    return err
}

// StoreURL stores a short URL and its corresponding original URL in the database.
func StoreURL(db *sql.DB, shortURL, originalURL string) error {
	query := `INSERT INTO urls (short_url, original_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortURL, originalURL)
	return err
}

// GetOriginalURL retrieves the original URL associated with a given short URL from the database.
func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM urls WHERE short_url = ? LIMIT 1`
	err := db.QueryRow(query, shortURL).Scan(&originalURL)
	if err != nil {
        return "", err
    }
	return originalURL, nil
}

// URLPair represents a pair of short and original URLs
type URLPair struct {
    ShortURL    string
    OriginalURL string
}

// GetAllURLs retrieves all shortened URLs from the database
func GetAllURLs(db *sql.DB) ([]URLPair, error) {
    query := `SELECT short_url, original_url FROM urls ORDER BY id DESC LIMIT 10`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var urls []URLPair
    for rows.Next() {
        var pair URLPair
        if err := rows.Scan(&pair.ShortURL, &pair.OriginalURL); err != nil {
            return nil, err
        }
        urls = append(urls, pair)
    }
    return urls, nil
}