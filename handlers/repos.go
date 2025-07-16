package handlers

import (
	"GitLocker/git"
	"encoding/json"
	"net/http"
	"os"
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

func GetRepo(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(storageRoot)
	if err != nil {
		http.Error(w, "failed to read repos", http.StatusInternalServerError)
		return
	}

	var repos []string
	for _, f := range files {
		if f.IsDir() {
			repos = append(repos, f.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repos)
}

func GetRepoDetails(w http.ResponseWriter, r *http.Request) {
	// URL: /api/repo-details/{name}
	// Extract repo name from the path
	prefix := "/api/repo-details/"
	if len(r.URL.Path) <= len(prefix) {
		http.Error(w, "missing repo name", http.StatusBadRequest)
		return
	}
	repoName := r.URL.Path[len(prefix):]
	if repoName == "" {
		http.Error(w, "missing repo name", http.StatusBadRequest)
		return
	}

	// Dummy data for now
	repoDetails := map[string]interface{}{
		"name":          repoName,
		"owner":         "mattkje",
		"description":   "A blazing fast self-hosted git service.",
		"stars":         42,
		"forks":         7,
		"watchers":      10,
		"updatedAt":     "2024-06-10",
		"defaultBranch": "main",
		"files": []map[string]string{
			{"name": "README.md", "type": "file"},
			{"name": "go.mod", "type": "file"},
			{"name": "cmd", "type": "dir"},
			{"name": "internal", "type": "dir"},
			{"name": "main.go", "type": "file"},
		},
		"lastCommit": map[string]string{
			"message": "Add authentication",
			"hash":    "d4e5f6g",
			"author":  "mattkje",
			"date":    "2024-06-11",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repoDetails)
}
