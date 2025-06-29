package silo

import (
	"fmt"
	"os"
	"path/filepath"
)

type Silo struct {
	Root    string
	DataDir string
}

func InitSilo(root string) error {
	info, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Can't create silo. Directory does not exist")
		}
		return fmt.Errorf("Failed to access directory %s: %w", root, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("Can't create silo here. %s is not a directory", root)
	}

	// make sure silo does not already exist there
	siloPath := filepath.Join(root, ".silo")
	if _, err := os.Stat(siloPath); !os.IsNotExist(err) {
		return fmt.Errorf("Silo already exists at %s", siloPath)
	}

	os.MkdirAll(siloPath, 0755)
	return nil
}

func FindEffectiveSilo(maybeRoot *string) (Silo, error) {
	if maybeRoot != nil && *maybeRoot != "" {
		if !IsSilo(*maybeRoot) {
			return Silo{}, fmt.Errorf("the specified path is not a valid silo root: %s", *maybeRoot)
		}
		return Silo{Root: *maybeRoot, DataDir: filepath.Join(*maybeRoot, ".silo")}, nil
	}

	siloRoot, err := FindRootFromCWD()
	if err != nil {
		return Silo{}, fmt.Errorf("failed to find silo root: %w", err)
	}
	return Silo{Root: siloRoot, DataDir: filepath.Join(siloRoot, ".silo")}, nil
}

func IsSilo(path string) bool {
	// path is a directory and it contains a .silo subdirectory
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		siloPath := filepath.Join(path, ".silo")
		if info, err := os.Stat(siloPath); err == nil && info.IsDir() {
			return true
		}
	}
	return false
}

func FindRootFromCWD() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}
	return FindRoot(currentDir)
}

func FindRoot(startPath string) (string, error) {
	dir := startPath

	for {
		siloPath := filepath.Join(dir, ".silo")
		info, err := os.Stat(siloPath)
		if err == nil && info.IsDir() {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf(".silo directory not found")
}
