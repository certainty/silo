package silo

import (
	"fmt"
	"os"
	"path/filepath"
)

type Silo struct {
	Root string
}

func FindRootFromCurrent() (Silo, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return Silo{}, fmt.Errorf("failed to get current working directory: %w", err)
	}
	return FindRoot(currentDir)
}

func FindRoot(startPath string) (Silo, error) {
	dir := filepath.Dir(startPath)
	for {
		siloPath := filepath.Join(dir, ".silo")
		info, err := os.Stat(siloPath)
		if err == nil && info.IsDir() {
			return Silo{Root: dir}, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return Silo{}, fmt.Errorf(".silo directory not found")
}

func (s Silo) TagsDir() string {
	return filepath.Join(s.Root, ".silo", "tags")
}
