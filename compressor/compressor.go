package compressor

import (
	"log"

	"github.com/h2non/bimg"
)

func Compress(file File) File {
	return ImageCompress(file)
}

func ImageCompress(inputFile File) File {
	outputFile := File{}
	outputFile.FileName = inputFile.FileName
	var err error
	outputFile.File, err = bimg.NewImage(inputFile.File).Process(bimg.Options{
		Quality:       80,   // Set quality (e.g., 80)
		StripMetadata: true, // Remove unnecessary metadata
	})

	if err != nil {
		log.Println("Image Compress Failed")
		return inputFile
	}

	return outputFile
}
