package main

import (
	"os"

	"github.com/4004-prison/cli"
	"github.com/czlingo/wormhole/cmd"
)

func main() {
	app := &cli.App{
		Name:        "Wormhole",
		Usage:       "wormhole command",
		Description: "",
		Commands: []*cli.Command{
			cmd.CMDVersion,
			cmd.CMDServer,
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
