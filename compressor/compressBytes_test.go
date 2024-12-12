package compressor

import (
	"reflect"
	"testing"
)

func TestCompressBytes(t *testing.T) {

	tests := []struct {
		name      string
		inputData string
		wantErr   bool
	}{
		{name: "successful 1", inputData: "hi", wantErr: false},
		{name: "successful 2", inputData: "Ultimate Packer for eXecutables", wantErr: false},
		{
			name: "successful 3",
			inputData: ` 
			Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2024
			UPX 4.2.2       Markus Oberhumer, Laszlo Molnar & John Reiser    Jan 3rd 2024`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBytes := []byte(tt.inputData)
			got, err := CompressBytes(inputBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompressBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			decompressed, err := DecompressBytes(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecompressBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			outputData := string(decompressed)

			if !reflect.DeepEqual(tt.inputData, outputData) {
				t.Errorf("CompressBytes() = %v, want %v", got, decompressed)
			}
		})
	}
}
