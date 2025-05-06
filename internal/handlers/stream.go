package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	songName := strings.TrimPrefix(r.URL.Path, "/stream/")
	filePath := filepath.Join("static", "songs", songName)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		http.Error(w, "Error reading file metadata", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "inline")
	http.ServeContent(w, r, songName, stat.ModTime(), file)
}