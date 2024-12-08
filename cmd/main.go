package main

import (
	"log"
	"os"

	"github.com/sumit-behera-in/go_cli_template/cmds"
	"github.com/urfave/cli/v2"
)

const (
	version = "1.2.1-template"
)

func main() {
	app := &cli.App{
		Name:    "go_cli_template",
		Usage:   "This is a template used to create cli applications",
		Version: version,
		Commands: []*cli.Command{
			cmds.Command(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "flag",
				Usage:    "Global flag",
				Value:    "default value",
				Required: true,
			},
		},
		Before: func(ctx *cli.Context) error {
			// what to do before app execution
			return nil
		},
		Action: func(ctx *cli.Context) error {
			return nil
		},
		After: func(ctx *cli.Context) error {
			// what to do after app execution
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
