
# GoCompress

![GoCompress Logo](https://via.placeholder.com/600x200.png?text=GoCompress) <!-- Replace with your logo -->

GoCompress is a powerful CLI tool for file compression built using **Golang** and **Cobra**. It supports multiple compression algorithms, including **Huffman coding** and **Run-Length Encoding (RLE)**, providing efficient and easy-to-use file management for your projects.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Compressing Files](#compressing-files)
  - [Decompressing Files](#decompressing-files)
- [Algorithms](#algorithms)
  - [Huffman Coding](#huffman-coding)
  - [Run-Length Encoding (RLE)](#run-length-encoding-rle)
- [Contributing](#contributing)
- [License](#license)

## Features
- **Multiple Compression Algorithms**: Compress and decompress files using Huffman coding and RLE.
- **Fast and Efficient**: Leverage the speed and performance of Golang for high-quality compression.
- **User-Friendly CLI**: Built with Cobra for easy command-line navigation.

## Installation

To install GoCompress, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/GoCompress.git
   cd GoCompress
   ```

2. Build the application:
   ```bash
   go build -o fcompressor
   ```

3. Run the application:
   ```bash
   ./fcompressor
   ```

## Usage

### Compressing Files
To compress a file, use the following command:
```bash
./fcompressor rle -i input.txt -o output.rle -a compress
```

### Decompressing Files
To decompress a file, use:
```bash
./fcompressor rle -i output.rle -o decompressed.txt -a decompress
```

## Algorithms

### Huffman Coding
Huffman coding is a widely used method for lossless data compression. It uses variable-length codes to represent characters based on their frequencies, allowing more frequent characters to be represented with shorter codes.

### Run-Length Encoding (RLE)
RLE is a simple compression algorithm that replaces sequences of repeated characters with a single character and its count. For example, the string "AAAABBBCCDAA" becomes "4A3B2C1D2A".

## Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch and open a pull request.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to reach out if you have any questions or need assistance!
