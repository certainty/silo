package tags

import (
	"fmt"
	"github.com/certainty/silo/internal/silo"
	"os"
	"path/filepath"
)

type Manager struct {
	Silo silo.Silo
}

func (m *Manager) AddTag(filePath, tag string) error {
	absFile, err := filepath.Abs(filePath)

	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	tagDir := filepath.Join(m.Silo.TagsDir(), filepath.FromSlash(tag))
	if err := os.MkdirAll(tagDir, 0755); err != nil {
		return fmt.Errorf("failed to create tag directory: %w", err)
	}

	relPath, err := filepath.Rel(tagDir, absFile)
	if err != nil {
		return fmt.Errorf("failed to get relative path: %w", err)
	}

	linkName := filepath.Join(tagDir, filepath.Base(absFile))
	if err := os.Symlink(relPath, linkName); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	return nil
}
