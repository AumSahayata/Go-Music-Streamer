package main

import (
	"log"
	"net/http"

	"music-streamer/internal/db"
	"music-streamer/internal/utils"
	"music-streamer/internal/handlers"
)

func main() {
	database := db.Init()
	utils.SeedSongs(database, "static/songs")

	http.HandleFunc("/stream/", handlers.StreamHandler)
	http.HandleFunc("/upload", handlers.UploadHandler(database))
	http.HandleFunc("/upload-form", handlers.UploadFormHandler)
	http.HandleFunc("/api/search", handlers.SearchSongsHandler(database))


	http.HandleFunc("/", handlers.IndexHandler(database))
	
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
