package compressor

import (
	"log"
	"path/filepath"

	"github.com/h2non/bimg"
)

func findType(ext string) string {

	switch ext {
	case ".jpg", ".jpeg", ".png", ".webp", ".bmp":
		return imageFileType
	}

	return ""
}

func Compress(file File) File {
	fileExtension := filepath.Ext(file.FileName)

	fileType := findType(fileExtension)

	if fileType == imageFileType {
		return ImageCompress(file)
	}

	return file
}

func ImageCompress(inputFile File) File {

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
		return inputFile
	}

	return outputFile
}
