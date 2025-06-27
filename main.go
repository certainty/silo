package main

import (
	"github.com/alecthomas/kong"
	"github.com/certainty/silo/cmd"
)

type CLI struct {
	Init cmd.InitCmd `cmd:"" name:"init" help:"Initialize a new silo in current directory or specified path"`
	Tags cmd.TagsCmd `cmd:"" name:"tags" aliases:"t" help:"Manage tags"`
}

func main() {
	var cli CLI
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
