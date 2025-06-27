package tags

import (
	"fmt"
)

type TagsCmd struct {
	Add AddCmd `cmd:"" name:"add" aliases:"a" help:"Add a tag to a file"`
}

type AddCmd struct {
	File string `arg:"" help:"File to tag"`
	Tag  string `arg:"" help:"Tag to add"`
}

func (cmd *AddCmd) Run() error {
	fmt.Println("Managing tags for file:", cmd.File)
	return nil
}
