package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/brkcnr/url-shortener/internal/controllers"
	"github.com/brkcnr/url-shortener/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

// main initializes the SQLite database, creates the urls table if it doesn't exist, and sets up the HTTP routes for the URL shortener.
func main() {
	slite, err := sql.Open("sqlite3", "db.sqlite")
	if err!= nil {
        log.Printf("Failed to open database: %v", err)
        log.Fatal(err)
    }
	defer slite.Close()

	if err := db.CreateTable(slite); err != nil {
		log.Printf("Failed to create table: %v", err)
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))

	// Create middleware to inject db into context
	dbMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", slite)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}

	http.HandleFunc("/shorten", controllers.Shorten(slite))

	http.HandleFunc("/", dbMiddleware(func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if path == "/" {
			controllers.ShowIndex(writer, request)
			return
		}
		controllers.Proxy(slite)(writer, request)
	}))

	http.HandleFunc("/history", dbMiddleware(controllers.ToggleHistory))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
