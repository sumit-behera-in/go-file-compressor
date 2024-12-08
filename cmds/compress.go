package cmds

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sumit-behera-in/go_file_compressor/compressor"
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
			fPath := ctx.String("file")
			outputPath := ctx.String("out")
			file := compressor.File{}

			file.FileName = filepath.Base(fPath)
			var err error
			if file.File, err = os.ReadFile(fPath); err != nil {
				return errors.New("the file in the file path not found")
			}
			file = compressor.Compress(file)
			if outputPath == "" {
				outputPath, err = GetDefaultDownloadPath()
				if err != nil {
					return err
				}
			}

			outputPath += string(filepath.Separator) + "compressed_" + file.FileName

			err = os.WriteFile(outputPath, file.File, 0666)
			if err != nil {
				return err
			}
			fmt.Printf("File %s successfully compressed to %s\n", file.FileName, outputPath)
			return nil
		},
	}
}
