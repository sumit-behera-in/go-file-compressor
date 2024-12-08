package cmds

import (
	"github.com/urfave/cli/v2"
)

func Compress() *cli.Command {
	return &cli.Command{
		Name:        "Compress",
		Usage:       "Compresses the input file",
		Description: "This command accepts the file location using the --file flag and the output location using the --out flag from the user to compress the file supplied by the input and store it into the output directory. ",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Usage:    "Absolute path of the input file to be compressed",
				Required: true,
			}, &cli.StringFlag{
				Name:     "out",
				Usage:    "Absolute path of the output path where the file will be stored after being compressed",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}
}