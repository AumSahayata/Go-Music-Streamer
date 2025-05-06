package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

type Song struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Filename string `json:"filename"`
}

func SearchSongsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request)  {
		query := strings.ToLower(r.URL.Query().Get("q"))

		var rows *sql.Rows
		var err error

		if query == "" {
			rows, err = db.Query("SELECT id, title, artist, filename FROM songs")
		}else {
			query = "%" + query + "%"
			rows, err = db.Query(`
				SELECT id, title, artist, filename FROM songs WHERE LOWER(title) LIKE ?
				OR LOWER(artist) LIKE ?;
			`, query, query)
		}

		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var songs []Song
		for rows.Next() {
			var s Song
			if err := rows.Scan(&s.ID, &s.Title, &s.Artist, &s.Filename); err == nil {
				songs = append(songs, s)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)
	}
}