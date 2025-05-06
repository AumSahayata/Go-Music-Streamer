package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite", "./songs.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS songs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		artist TEXT,
		filename TEXT UNIQUE
	);`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	return db
}
