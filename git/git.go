package git

import (
	"fmt"
	"os"
	"os/exec"
)

func InitBareRepo(name string) error {
	path := fmt.Sprintf("./repos/%s.git", name)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("repo already exists")
	}
	return exec.Command("git", "init", "--bare", path).Run()
}
