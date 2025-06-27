package silo

import (
	"os"
	"path/filepath"

	"github.com/certainty/silo/internal/ux"
)

type InitCmd struct {
	Path string `arg:"" optional:"" help:"Path to initialize the silo. Defaults to current directory."`
}

func (cmd *InitCmd) Run() error {
	if cmd.Path == "" {
		cmd.Path = "."
	}

	if _, err := os.Stat(filepath.Join(cmd.Path, ".silo")); !os.IsNotExist(err) {
		ux.Info("Silo already initialized at %s", cmd.Path)
		return nil
	}

	os.MkdirAll(filepath.Join(cmd.Path, ".silo"), 0755)
	// TODO: setup initial configuration files or directories

	ux.Info("Silo initialized at %s", cmd.Path)
	return nil
}
