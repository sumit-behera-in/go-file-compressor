# go_file_compressor
A Go library for compressing images, videos, and documents.

## Dependencies
### Runtime Dependencies 
- <a href="https://ffmpeg.org/">ffmpeg</a> – Required for video compression
### Buildtime Dependencies
- <a href="https://upx.github.io/">upx</a> – For executable file compression
- <a href="https://www.gnu.org/software/make/">make</a> – For building the project

## Features
- **Compress(file File)**: Compresses any input file and returns the most optimized version, saved as a `.gz` file.
- **ImageCompress(inputFile File)**: Compresses an image file and returns the compressed result in the same format.
- **VideoCompress(inputFile File)**: Compresses a video file and returns the optimized compressed video.
- **CompressBytes(inputData []byte)**: Compresses byte data and returns the result as a `.gz` compressed file.
- **DecompressBytes(compressedData []byte)**: Decompresses the provided `.gz` compressed byte data and returns the original file.

## Contributing 💖
We welcome contributions to help improve and expand this library! Whether you're fixing bugs, adding new features, or enhancing the documentation, your contributions are highly appreciated. Made with love by the community, for the community.
