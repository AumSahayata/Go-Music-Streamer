package utils

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

func SeedSongs(db *sql.DB, songDir string) {
	dirEntries, err := os.ReadDir(songDir)
	if err != nil {
		log.Println("Could not read songs directory:", err)
		return
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".mp3" {
			fullPath := filepath.Join(songDir, entry.Name())
			err := InsertSongFromFile(db, fullPath, entry.Name())
			if err != nil {
				log.Println("Failed to insert song from file:", entry.Name(), err)
			}
		}
	}
}
