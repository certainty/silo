package tags

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/certainty/silo/internal/silo"
	"github.com/ktr0731/go-fuzzyfinder"
)

type manager struct {
	Silo    silo.Silo
	TagsDir string
}

func NewManager(s silo.Silo) *manager {
	return &manager{
		Silo:    s,
		TagsDir: filepath.Join(s.DataDir, "tags"),
	}
}

func (m *manager) InitTagsDir() error {
	if _, err := os.Stat(m.TagsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(m.TagsDir, 0755); err != nil {
			return fmt.Errorf("failed to create tags directory: %w", err)
		}
	}
	return nil
}

func (m *manager) AddTags(filePath string, tags []string) error {
	if err := m.InitTagsDir(); err != nil {
		return fmt.Errorf("failed to initialize tags directory: %w", err)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	for _, tag := range tags {
		if err := m.addTag(filePath, tag); err != nil {
			return fmt.Errorf("failed to add tag %s to file %s: %w", tag, filePath, err)
		}
	}
	return nil
}

func (m *manager) AddTagsInteractively(filePath string) error {
	if err := m.InitTagsDir(); err != nil {
		return fmt.Errorf("failed to initialize tags directory: %w", err)
	}

	// make sure the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	// ask the user for tags using bubbletea
	tags, err := m.AllTags()
	if err != nil {
		return err
	}

	idxs, err := fuzzyfinder.FindMulti(tags,
		func(i int) string {
			return tags[i]
		},
		fuzzyfinder.WithPromptString("Select tags > "),
	)
	if err != nil {
		return err
	}

	var selected []string
	for _, i := range idxs {
		selected = append(selected, tags[i])
	}

	if len(selected) != 0 {
		return m.AddTags(filePath, selected)
	}
	return nil
}

func (m *manager) addTag(filePath, tag string) error {
	absFile, err := filepath.Abs(filePath)

	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	tagDir := filepath.Join(m.TagsDir, filepath.FromSlash(tag))
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

func (m *manager) AllTags() ([]string, error) {
	var tags []string

	err := filepath.Walk(m.TagsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == m.TagsDir {
			return nil
		}

		if info.IsDir() {
			relPath, err := filepath.Rel(m.TagsDir, path)
			if err != nil {
				return err
			}

			tag := strings.ReplaceAll(relPath, string(filepath.Separator), "/")
			tags = append(tags, tag)
			tags = append(tags, filepath.Base(path))
		}

		return nil
	})

	uniqueTags := removeDuplicates(tags)
	sort.Strings(uniqueTags)

	return uniqueTags, err
}

func removeDuplicates(input []string) []string {
	seen := make(map[string]struct{})
	var result []string

	for _, str := range input {
		if _, ok := seen[str]; !ok {
			seen[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}
