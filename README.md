# go_file_compressor
A Go library for compressing images, videos, and documents, featuring a CLI interface for seamless interaction. If no output path is specified, the compressed file is automatically saved to the default download folder of the operating system.

## Dependencies
### Runtime Dependencies
- <a href="https://ffmpeg.org/">ffmpeg</a> – Required for video compression.
### Buildtime Dependencies
- <a href="https://upx.github.io/">upx</a> – For compressing executable files.
- <a href="https://www.gnu.org/software/make/">make</a> – For building the project.

## Features
- **Compress(file File)**: Compresses any input file and returns the most optimized version, saved as a `.gz` file.
- **ImageCompress(inputFile File)**: Compresses an image file, returning the compressed result in the same image format.
- **VideoCompress(inputFile File)**: Compresses a video file, returning an optimized, compressed video.
- **CompressBytes(inputData []byte)**: Compresses byte data and returns the result as a `.gz` compressed file.
- **DecompressBytes(compressedData []byte)**: Decompresses `.gz` compressed byte data and returns the original file.

## CLI Commands
- `goFileCompressor Compress --file="inputFilePath" [--out="outputFilePath"]`  
  Compresses the file at `inputFilePath`. If no `--out` path is provided, the result will be saved to the default download folder.
  
- `goFileCompressor Decompress --file="inputFilePath" [--out="outputFilePath"]`  
  Decompresses the `.gz` file at `inputFilePath`. If no `--out` path is specified, the decompressed file will be saved to the default download folder.

## Contributing 💖
We encourage contributions to help improve and expand this library! Whether you're fixing bugs, adding new features, or improving documentation, your input is greatly appreciated. Made with love by the community, for the community.