package cmd

import (
	"github.com/alecthomas/kong"
	"github.com/certainty/silo/internal/silo"
	"github.com/certainty/silo/internal/ux"
)

type InitCmd struct{}

type InfoCmd struct{}

type CLI struct {
	Root *string `help:"Root of the silo. All operations are relative to this path." env:"SILO_ROOT" global:""`
	Info InfoCmd `cmd:"" name:"info" help:"Display information about the current silo"`
	Init InitCmd `cmd:"" name:"init" help:"Initialize a new silo in current directory or specified path"`
	Tags TagsCmd `cmd:"" name:"tags" aliases:"t" help:"Manage tags"`
}

func (cmd *InitCmd) Run(ctx *kong.Context, cli *CLI) error {
	path := "."
	if cli.Root != nil && *cli.Root != "" {
		path = *cli.Root
	}

	if err := silo.InitSilo(path); err != nil {
		return err
	}

	ux.Info("Silo initialized at %s", path)
	return nil
}

func (cmd *InfoCmd) Run(ctx *kong.Context, cli *CLI) error {
	s, err := silo.FindEffectiveSilo(cli.Root)
	if err != nil {
		return err
	}

	ux.Info("Silo root: %s", s.Root)
	return nil
}
