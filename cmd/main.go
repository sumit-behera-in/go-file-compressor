package main

import (
	"log"
	"os"

	"github.com/sumit-behera-in/go-file-compressor/cmds"
	"github.com/urfave/cli/v2"
)

const (
	version = "1.0.3"
)

func main() {
	app := &cli.App{
		Name:    "go-file-compressor",
		Usage:   "This is a go library used to compress images, videos and documents.",
		Version: version,
		Commands: []*cli.Command{
			cmds.Compress(),
			cmds.DeCompress(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
