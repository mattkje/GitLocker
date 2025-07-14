package handlers

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&creds)

	hash, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	_, err := h.DB.Exec("INSERT INTO users(username, password) VALUES(?, ?)", creds.Username, hash)
	if err != nil {
		http.Error(w, "User exists", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&creds)

	row := h.DB.QueryRow("SELECT password FROM users WHERE username = ?", creds.Username)
	var hash string
	if err := row.Scan(&hash); err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(creds.Password)) != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Login OK"))
}
