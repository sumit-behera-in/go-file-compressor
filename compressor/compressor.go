package compressor

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/bimg"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func Compress(file File) (File, error) {
	fileExtension := filepath.Ext(file.FileName)

	switch fileExtension {
	case ".jpg", ".jpeg", ".png", ".webp", ".bmp":
		return ImageCompress(file)
	case ".mp4", ".avi", ".mov", ".mkv", ".webm", ".flv", ".ogg", ".3gp", ".wmv", ".mpg":
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

	if outputFile.File, err = CompressBytes(outputFile.File); err != nil {
		log.Println("Image Compress to .gz Failed")
		return outputFile, err
	}

	return outputFile, nil
}

func VideoCompress(inputFile File) (File, error) {

	fileExt := filepath.Ext(inputFile.FileName)

	// Create the temp directory with the right permissions (0777 allows full permissions)
	err := os.MkdirAll(tempVideoDir, 0777)
	if err != nil {
		return inputFile, fmt.Errorf("failed to create temp directory: %v", err)
	}

	// Write input to a temp file
	inputPath := tempVideoDir + inputFile.FileName
	if err := os.WriteFile(inputPath, inputFile.File, 0666); err != nil {
		return inputFile, fmt.Errorf("failed to write temp input file: %v", err)
	}

	// Define output file path
	outputPath := tempVideoDir + inputFile.FileName + "compressed" + fileExt

	// Run FFmpeg command
	err = ffmpeg_go.Input(inputPath).
		Output(outputPath, ffmpeg_go.KwArgs{
			"f":          fileExt[1:],
			"err_detect": "ignore_err",
			"fflags":     "+igndts",
			"pix_fmt":    "yuv420p",
		}).
		OverWriteOutput().
		Run()

	if err != nil {
		return inputFile, fmt.Errorf("ffmpeg execution failed: %v", err)
	}

	// Read the output file back into memory
	outputData, err := os.ReadFile(outputPath)
	if err != nil {
		return inputFile, fmt.Errorf("failed to read temp output file: %v", err)
	}

	// Clean up temp files
	if err = os.Chmod(inputPath, 0666); err != nil {
		panic(err.Error())
	}

	if err = os.Chmod(outputPath, 0666); err != nil {
		panic(err.Error())
	}

	if err = os.Remove(inputPath); err != nil {
		panic(err.Error())
	}
	if err = os.Remove(outputPath); err != nil {
		panic(err.Error())
	}

	if outputData, err = CompressBytes(outputData); err != nil {
		log.Println("Image Compress to .gz Failed")
		return File{FileName: inputFile.FileName, File: outputData}, err
	}

	// Return the compressed video
	return File{FileName: inputFile.FileName, File: outputData}, nil
}
