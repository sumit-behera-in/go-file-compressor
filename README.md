# go_file_compressor
A Go library for compressing images, videos, and documents, featuring a user-friendly CLI interface for seamless interaction. If no output path is provided, the compressed file will be saved to the default download folder of your operating system.

## Features
- **Compress(file File)**: Compress any input file and return the most optimized version, saved as a `.gz` file.
- **ImageCompress(inputFile File)**: Compress image files, returning the compressed result in the same image format.
- **VideoCompress(inputFile File)**: Compress video files, optimizing them for size while retaining quality.
- **CompressBytes(inputData []byte)**: Compress byte data and return the result as a `.gz` compressed file.
- **DecompressBytes(compressedData []byte)**: Decompress `.gz` compressed byte data and return the original file.

## CLI Commands
- `goFileCompressor Compress --file="inputFilePath" [--out="outputFilePath"]`  
  Compress the file at `inputFilePath`. If no output path (`--out`) is provided, the result will be saved in the default download folder.

- `goFileCompressor Decompress --file="inputFilePath" [--out="outputFilePath"]`  
  Decompress the `.gz` file at `inputFilePath`. If no output path (`--out`) is specified, the decompressed file will be saved in the default download folder.

## Dependencies
### Runtime Dependencies
- <a href="https://ffmpeg.org/">ffmpeg</a> – Required for video compression.

### Buildtime Dependencies
- <a href="https://upx.github.io/">upx</a> – For executable file compression.
- <a href="https://www.gnu.org/software/make/">make</a> – For building the project.

## Disclaimer
During video compression, both the original and the compressed video files are temporarily stored in the temp/ directory. Once the compression process is complete, these files are deleted. To avoid conflicts and potential data loss, please refrain from manually interacting with this directory during compression, as it may overwrite existing files.

## Contributing 💖
We welcome contributions to help improve and expand this library! Whether you're fixing bugs, adding new features, or improving the documentation, your help is greatly appreciated. Made with love by the community, for the community.