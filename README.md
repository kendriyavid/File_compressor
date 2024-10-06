
# GoCompress

GoCompress is a simple and efficient file compression tool implemented in Go using Huffman encoding. It allows users to compress and decompress files seamlessly.

## Features

- **Huffman Encoding**: Utilizes Huffman coding for effective lossless data compression.
- **File Compression**: Compresses input files into smaller sizes.
- **File Decompression**: Restores original files from compressed formats.
- **Tree Structure Storage**: Encodes the Huffman tree structure within the compressed file for accurate decompression.

## Installation

To get started with GoCompress, ensure you have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/GoCompress.git
   cd GoCompress
   ```

2. Build the project:

   ```bash
   go build -o gocompress
   ```

## Usage

### Compressing a File

To compress a file, use the following command:

```bash
./gocompress compress <input_file> <output_file>
```

### Decompressing a File

To decompress a file, use the following command:

```bash
./gocompress decompress <input_file> <output_file>
```

### Example

```bash
# Compressing a file
./gocompress compress example.txt example.txt.gz

# Decompressing a file
./gocompress decompress example.txt.gz decompressed_example.txt
```

## How It Works

1. **Build Frequency Table**: Reads the input file and creates a frequency table of the characters.
2. **Build Huffman Tree**: Constructs a binary tree based on character frequencies.
3. **Generate Codes**: Assigns binary codes to each character based on their position in the tree.
4. **Compress**: Writes the Huffman tree and the encoded data to the output file.
5. **Decompress**: Reads the Huffman tree from the input file and decodes the compressed data.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature/YourFeature`).
6. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries, feel free to contact me at [your-email@example.com](mailto:your-email@example.com).
