package compressor

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/h2non/bimg"
)

func Compress(file File) (File, error) {
	fileExtension := filepath.Ext(file.FileName)

	switch fileExtension {
	case ".jpg", ".jpeg", ".png", ".webp", ".bmp":
		return ImageCompress(file)
	case ".mp4":
		return VideoCompress(file)
	}

	return file, errors.New("unsupported file format")
}

func ImageCompress(inputFile File) (File, error) {

	outputFile := File{}
	outputFile.FileName = inputFile.FileName
	var err error
	outputFile.File, err = bimg.NewImage(inputFile.File).Process(bimg.Options{
		Type:          bimg.WEBP,
		Quality:       75,   // Set quality (e.g., 80)
		StripMetadata: true, // Remove unnecessary metadata
	})

	if err != nil {
		log.Println("Image Compress Failed")
		return inputFile, err
	}

	return outputFile, nil
}

func VideoCompress(inputFile File) (File, error) {

	outputFile := File{}
	outputFile.FileName = inputFile.FileName
	var err error
	outputFile.File, err = bimg.NewImage(inputFile.File).Process(bimg.Options{
		Type:          bimg.WEBP,
		Quality:       75,   // Set quality (e.g., 80)
		StripMetadata: true, // Remove unnecessary metadata
	})

	if err != nil {
		log.Println("Image Compress Failed")
		return inputFile, err
	}

	return outputFile, nil
}
