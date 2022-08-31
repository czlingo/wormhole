package cmd

import (
	"fmt"

	"github.com/4004-prison/cli"
	"github.com/czlingo/wormhole/config"
)

var CMDVersion = &cli.Command{
	Name: "version",
	Action: func(ctx *cli.Context) {
		fmt.Printf("Wormhole version %s\n", config.VERSION)
	},
}
