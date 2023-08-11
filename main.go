package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		panic(err)
	}

	_, err = worktree.Add("main.go")
	if err != nil {
		panic(err)
	}

	_, err = worktree.Commit("testing", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "mlctrez",
			Email: "mlctrez@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		panic(err)
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(filepath.Join(dir, ".github_token"))
	if err != nil {
		panic(err)
	}
	token := strings.TrimSpace(string(file))

	fmt.Println(token)

	err = repo.Push(&git.PushOptions{
		Auth: &http.BasicAuth{Username: token},
	})
	if err != nil {
		panic(err)
	}

}
