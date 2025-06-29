package cmd

import (
	"github.com/alecthomas/kong"
	"github.com/certainty/silo/internal/silo"
	"github.com/certainty/silo/internal/tags"
)

type TagsCmd struct {
	Add AddCmd `cmd:"" name:"add" aliases:"a" help:"Add a tag to a file"`
}

type AddCmd struct {
	File string   `arg:"" help:"File to tag"`
	Tags []string `arg:"" optional:"" help:"Tags to add (comma or space separated)"`
}

func (cmd *AddCmd) Run(ctx *kong.Context, cli *CLI) error {
	s, err := silo.FindEffectiveSilo(cli.Root)
	if err != nil {
		return err
	}
	m := tags.NewManager(s)

	if cmd.Tags == nil || len(cmd.Tags) == 0 {
		if err := m.AddTagsInteractively(cmd.File); err != nil {
			return err
		}
	} else {
		if err := m.AddTags(cmd.File, cmd.Tags); err != nil {
			return err
		}
	}
	return nil
}
