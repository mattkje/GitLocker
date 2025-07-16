package main

import (
	"GitLocker/handlers"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	auth := &handlers.AuthHandler{DB: db}
	http.HandleFunc("/api/register", auth.Register)
	http.HandleFunc("/api/login", auth.Login)
	http.HandleFunc("/api/create-repo", handlers.CreateRepo)
	http.HandleFunc("/api/repos/", handlers.GetRepo)
	http.HandleFunc("/api/repo-details/", handlers.GetRepoDetails)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/assets/", fs) // serve static assets

	// Fallback handler for SPA routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If the file exists, serve it
		path := filepath.Join("./frontend/dist", r.URL.Path)
		if _, err := os.Stat(path); err == nil {
			http.ServeFile(w, r, path)
			return
		}
		// Otherwise, serve index.html
		http.ServeFile(w, r, "./frontend/dist/index.html")
	})

	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
