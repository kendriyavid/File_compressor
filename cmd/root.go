package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fcompressor",
	Short: "A file compressor CLI application using multiple algorithms",
	Long: `fcompressor is a CLI tool to compress and decompress files using various algorithms 
like Huffman, RLE, and LZW. This tool supports easy and fast file compression for your projects.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
