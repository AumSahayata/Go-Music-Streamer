package utils

import (
	"database/sql"
	"os"

	"github.com/dhowden/tag"
)

func InsertSongFromFile(db *sql.DB, filePath string, filename string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	title := filename
	artist := "Unknown"

	if metadata, err := tag.ReadFrom(file); err == nil {
		if metadata.Title() != "" {
			title = metadata.Title()
		}
		if metadata.Artist() != "" {
			artist = metadata.Artist()
		}
	}

	_, err = db.Exec(
		"INSERT OR IGNORE INTO songs (title, artist, filename) VALUES (?, ?, ?)",
		title, artist, filename,
	)
	return err
}
