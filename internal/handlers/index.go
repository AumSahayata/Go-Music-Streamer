package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"music-streamer/internal/models"
)

func IndexHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		rows, err := db.Query("SELECT title, artist, filename FROM songs")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var songs []models.Song
		for rows.Next() {
			var s models.Song
			if err := rows.Scan(&s.Title, &s.Artist, &s.Filename); err == nil {
				songs = append(songs, s)
			}
		}

		tmpl.Execute(w, songs)
	}
}