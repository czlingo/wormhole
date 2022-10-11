package cmd

import (
	"github.com/4004-prison/cli"
	"github.com/czlingo/wormhole/internal/define"
	"github.com/czlingo/wormhole/internal/model"
	"github.com/czlingo/wormhole/internal/route"
)

var CMDServer = &cli.Command{
	Name:  "server",
	Usage: "server",
	Action: func(ctx *cli.Context) {
		port := ctx.String("--port")

		config := define.InitConfig()
		_, err := model.New(config)
		if err != nil {
			panic(err)
		}

		route.New().RunOrDie(port)
	},
	Flags: cli.FlagSet{
		&cli.String{
			Name:  "--port",
			Usage: "Server port. default 6789",
		},
	},
}
