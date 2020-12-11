//go:generate go run main.go tutorial.go
package main

import (
	"fmt"
	"github.com/golangee/i18n"
	"os"
	"path/filepath"
)

func main() {
	// invoke the generator in your current project. It will process the entire module.
	if err := i18n.Bundle(); err != nil {
		panic(err)
	}

	if err := createTutorialDetails(); err != nil {
		panic(err)
	}
}

func createTutorialDetails() error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	dir, err = modRoot(dir)
	if err != nil {
		return err
	}

	if err := buildIndex(dir, filepath.Join(dir, "static", "assets", "gen")); err != nil {
		return err
	}

	return nil
}

func modRoot(dir string) (string, error) {
	if dir == "/" || dir == "" {
		return "", fmt.Errorf("no go.mod found")
	}

	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err != nil {
		return modRoot(filepath.Dir(dir))
	}

	return dir, nil
}
