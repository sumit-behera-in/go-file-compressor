package cmds

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "Command",
		Usage: "Demo command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "flag",
				Usage:    "demo flag",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}
}
