package compressor

import (
	"os"
	"testing"
)

func getFile(t *testing.T) []byte {
	file, err := os.ReadFile("/home/sk/Desktop/go-file-compressor/bin/a.mp4")
	if err != nil {
		t.Error(err.Error())
	}
	return file
}

func TestVideoCompress(t *testing.T) {

	tests := []struct {
		name      string
		inputFile File
		wantErr   bool
	}{
		{
			name:      "compress success",
			inputFile: File{FileName: "a.mp4", File: getFile(t)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VideoCompress(tt.inputFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			err = os.WriteFile("bin/a1.mp4", got.File, 0666)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
