package main

import (
	"log"
	"os"

	"github.com/sumit-behera-in/go_file_compressor/cmds"
	"github.com/urfave/cli/v2"
)

const (
	version = "1.0.0"
)

func main() {
	app := &cli.App{
		Name:    "go_file_compressor",
		Usage:   "This is a go library used to compress images, videos and documents.",
		Version: version,
		Commands: []*cli.Command{
			cmds.Compress(),
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
