package main

import (
	"GitLocker/handlers"
	"log"
	"net/http"
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

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
