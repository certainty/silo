package cmd

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	"github.com/certainty/silo/internal/silo"
	"github.com/certainty/silo/internal/ux"
)

type InitCmd struct {
	Path string `arg:"" optional:"" help:"Path to initialize the silo. Defaults to current directory."`
}

type InfoCmd struct{}

type CLI struct {
	Root *string `help:"Root of the silo. All operations are relative to this path." env:"SILO_ROOT" global:""`
	Info InfoCmd `cmd:"" name:"info" help:"Display information about the current silo"`
	Init InitCmd `cmd:"" name:"init" help:"Initialize a new silo in current directory or specified path"`
	Tags TagsCmd `cmd:"" name:"tags" aliases:"t" help:"Manage tags"`
}

func (cmd *InitCmd) Run(ctx *kong.Context, cli *CLI) error {
	if cmd.Path == "" {
		cmd.Path = "."
	}

	if _, err := os.Stat(filepath.Join(cmd.Path, ".silo")); !os.IsNotExist(err) {
		ux.Info("Silo already initialized at %s", cmd.Path)
		return nil
	}

	os.MkdirAll(filepath.Join(cmd.Path, ".silo"), 0755)

	ux.Info("Silo initialized at %s", cmd.Path)
	return nil
}

func (cmd *InfoCmd) Run(ctx *kong.Context, cli *CLI) error {
	s, err := silo.FindEffectiveSilo(cli.Root)
	if err != nil {
		ux.Error("Failed to find silo root: %v", err)
		return err
	}

	ux.Info("Silo root: %s", s.Root)
	return nil
}
