package cmd

import (
	"fcompressor/rle" // Update with your actual package path
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rleInputFile  string
	rleOutputFile string
	rleAction     string // "compress" or "decompress"
)

var rleCmd = &cobra.Command{
	Use:   "rle",
	Short: "Compress or decompress files using RLE",
	Run: func(cmd *cobra.Command, args []string) {
		// Debugging statements
		fmt.Printf("Input file: %s\n", rleInputFile)
		fmt.Printf("Output file: %s\n", rleOutputFile)
		fmt.Printf("Action: %s\n", rleAction)

		if rleAction == "compress" {
			err := rle.CompressFile(rleInputFile, rleOutputFile)
			if err != nil {
				fmt.Printf("Error compressing file: %v\n", err)
				return
			}
			fmt.Printf("Successfully compressed '%s' to '%s'\n", rleInputFile, rleOutputFile)
		} else if rleAction == "decompress" {
			err := rle.DecompressFile(rleInputFile, rleOutputFile)
			if err != nil {
				fmt.Printf("Error decompressing file: %v\n", err)
				return
			}
			fmt.Printf("Successfully decompressed '%s' to '%s'\n", rleInputFile, rleOutputFile)
		} else {
			fmt.Println("Invalid action. Use 'compress' or 'decompress'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(rleCmd)

	// Define flags for the RLE command
	rleCmd.Flags().StringVarP(&rleInputFile, "input", "i", "", "Input file (required)")
	rleCmd.Flags().StringVarP(&rleOutputFile, "output", "o", "", "Output file (required)")
	rleCmd.Flags().StringVarP(&rleAction, "action", "a", "compress", "Action to perform: compress or decompress")

	// Mark required flags
	rleCmd.MarkFlagRequired("input")
	rleCmd.MarkFlagRequired("output")

	// Debugging statement
	fmt.Println("RLE command initialized with flags:")
	fmt.Printf("  Input: %s\n", rleInputFile)
	fmt.Printf("  Output: %s\n", rleOutputFile)
	fmt.Printf("  Action: %s\n", rleAction)
}
