package controllers

import (
	"database/sql"
	"net/http"
	"text/template"

	"github.com/brkcnr/url-shortener/internal/db"
)

// ShowIndex handles the rendering of the index page template.
func ShowIndex(w http.ResponseWriter, r *http.Request) {
	database := r.Context().Value("db").(*sql.DB)
	
	urls, err := db.GetAllURLs(database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		URLs []db.URLPair
	}{
		URLs: urls,
	}

	tmpl, err := template.ParseFiles("internal/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}