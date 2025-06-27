package cmd

import (
	"fmt"

	"github.com/certainty/silo/internal/silo"
	"github.com/certainty/silo/internal/tags"
)

type TagsCmd struct {
	Add AddCmd `cmd:"" name:"add" aliases:"a" help:"Add a tag to a file"`
}

type AddCmd struct {
	File string `arg:"" help:"File to tag"`
	Tag  string `arg:"" help:"Tag to add"`
}

func (cmd *AddCmd) Run() error {
	s, err := silo.FindRootFromCurrent()
	if err != nil {
		return err
	}
	m := tags.Manager{Silo: s}
	if err := m.AddTag(cmd.File, cmd.Tag); err != nil {
		return err
	}
	fmt.Printf("Tagged %s with %s\n", cmd.File, cmd.Tag)
	return nil

}
