package compressor

import (
	"bytes"
	"compress/gzip"
	"io"
)

// CompressBytes() compresses the input data and returns the compressed []byte.
// Always use the .gz extension at the end to avoid corruption.
func CompressBytes(inputData []byte) ([]byte, error) {
	var compressedBuffer bytes.Buffer

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(&compressedBuffer)
	defer gzipWriter.Close()

	// Write data to the gzip writer
	_, err := gzipWriter.Write(inputData)
	if err != nil {
		return inputData, err
	}

	// Close the writer to flush remaining data
	err = gzipWriter.Close()
	if err != nil {
		return inputData, err
	}

	return compressedBuffer.Bytes(), nil
}

// DecompressBytes() decompresses the input gzip []byte and returns the decompressed data.
// Provide a .gz file and receive the corresponding decompressed file as the result.
func DecompressBytes(compressedData []byte) ([]byte, error) {
	compressedBuffer := bytes.NewReader(compressedData)
	gzipReader, err := gzip.NewReader(compressedBuffer)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// Decompress the data
	var decompressedBuffer bytes.Buffer
	_, err = io.Copy(&decompressedBuffer, gzipReader)
	if err != nil {
		return compressedData, err
	}

	return decompressedBuffer.Bytes(), nil
}
