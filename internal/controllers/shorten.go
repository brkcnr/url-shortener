package controllers

import (
	"database/sql"
	"net/http"
	"strings"
	"text/template"

	"github.com/brkcnr/url-shortener/internal/db"
	"github.com/brkcnr/url-shortener/internal/url"
)

// Shorten handles the creation of a short URL for a given original URL.
func Shorten(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		originalURL := r.FormValue("url")
		if originalURL == "" {
			http.Error(w, "URL not provided", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
			originalURL = "https://" + originalURL
		}

		shortURL := url.Shorten(originalURL)

		if err := db.StoreURL(lite, shortURL, originalURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
		}

		data := map[string]string{
			"ShortURL": shortURL,
		}

		t, err := template.ParseFiles("internal/views/shorten.html")
		if err!= nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Proxy handles the redirection of a short URL to the original URL.
func Proxy(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Path[1:] // Remove leading slash
		if shortUrl == "" {
			http.Error(w, "URL not provided", http.StatusBadRequest)
			return
		}

		// Make sure the original URL includes the protocol
		origUrl, err := db.GetOriginalURL(lite, shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Ensure the URL has a protocol
		if !strings.HasPrefix(origUrl, "http://") && !strings.HasPrefix(origUrl, "https://") {
			origUrl = "https://" + origUrl
		}

		http.Redirect(w, r, origUrl, http.StatusTemporaryRedirect) // Changed to temporary redirect
	}
}