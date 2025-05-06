package handlers

import (
	"database/sql"
	"io"
	"music-streamer/internal/utils"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFormHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/upload.html")
}

func UploadHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
			return
		}

		file, header, err := r.FormFile("songFile")
		if err != nil {
			http.Error(w, "Failed to read file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		dstPath := filepath.Join("static/songs", header.Filename)
		out, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}

		defer out.Close()
		io.Copy(out, file)

		err = utils.InsertSongFromFile(db, dstPath, header.Filename)
		if err != nil {
			http.Error(w, "Failed to insert into DB", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}