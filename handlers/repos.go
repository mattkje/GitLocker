package handlers

import (
	"GitLocker/git"
	"encoding/json"
	"net/http"
)

type RepoRequest struct {
	Name string `json:"name"`
}

const storageRoot = "./repos"

func CreateRepo(w http.ResponseWriter, r *http.Request) {
	var req RepoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := git.InitBareRepo(storageRoot, req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Repo created"))
}
