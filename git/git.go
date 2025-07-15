package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func InitBareRepo(storageRoot, name string) error {
	// Validate repo name (simple example)
	if name == "" || len(name) > 100 {
		return fmt.Errorf("invalid repo name")
	}
	path := filepath.Join(storageRoot, fmt.Sprintf("%s.git", name))
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("repo already exists")
	}
	if err := os.MkdirAll(storageRoot, 0755); err != nil {
		return fmt.Errorf("failed to create storage root: %w", err)
	}
	return exec.Command("git", "init", "--bare", path).Run()
}
