package cmd

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/certainty/silo/internal/silo"
	"github.com/certainty/silo/internal/tags"
)

type TagsCmd struct {
	Add AddCmd `cmd:"" name:"add" aliases:"a" help:"Add a tag to a file"`
}

type AddCmd struct {
	File string   `arg:"" help:"File to tag"`
	Tags []string `arg:"" help:"Tags to add (comma or space separated)"`
}

func (cmd *AddCmd) Run(ctx *kong.Context, cli *CLI) error {
	s, err := silo.FindEffectiveSilo(cli.Root)
	if err != nil {
		return err
	}
	m := tags.Manager{Silo: s}
	if err := m.AddTags(cmd.File, cmd.Tags); err != nil {
		return err
	}
	fmt.Printf("Tagged %s with %s\n", cmd.File, cmd.Tags)
	return nil
}
