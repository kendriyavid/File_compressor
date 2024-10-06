package cmd

import (
	"fcompressor/huffman"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	huffmanInputFile  string
	huffmanOutputFile string
	isDecompress      bool
)

var huffmanCmd = &cobra.Command{
	Use:   "huffman",
	Short: "Compress or decompress files using Huffman coding",
	Long: `Use this command to either compress or decompress files using Huffman coding.
Provide the input file and output file, and use --decompress flag to decompress.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Check if input file exists
		if _, err := os.Stat(huffmanInputFile); os.IsNotExist(err) {
			return fmt.Errorf("input file '%s' does not exist", huffmanInputFile)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if isDecompress {
			err = huffman.DecompressFile(huffmanInputFile, huffmanOutputFile)
		} else {
			err = huffman.CompressFile(huffmanInputFile, huffmanOutputFile)
		}

		if err != nil {
			return fmt.Errorf("operation failed: %v", err)
		}

		action := "Compressed"
		if isDecompress {
			action = "Decompressed"
		}
		fmt.Printf("%s '%s' to '%s'\n", action, huffmanInputFile, huffmanOutputFile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(huffmanCmd)

	huffmanCmd.Flags().StringVarP(&huffmanInputFile, "input", "i", "", "Input file to compress or decompress")
	huffmanCmd.Flags().StringVarP(&huffmanOutputFile, "output", "o", "", "Output file to save the result")
	huffmanCmd.Flags().BoolVarP(&isDecompress, "decompress", "d", false, "Decompress the input file instead of compressing it")

	huffmanCmd.MarkFlagRequired("input")
	huffmanCmd.MarkFlagRequired("output")
}
