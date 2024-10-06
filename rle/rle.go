package rle

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Compress compresses the input string using RLE
func Compress(input string) string {
	var result strings.Builder
	length := len(input)

	i := 0
	for i < length {
		count := 1
		for i+1 < length && input[i] == input[i+1] {
			count++
			i++
		}
		result.WriteString(string(input[i]) + strconv.Itoa(count))
		i++
	}
	return result.String()
}

// Decompress decompresses the RLE encoded string
func Decompress(input string) string {
	var result strings.Builder
	length := len(input)

	i := 0
	for i < length {
		char := input[i]
		i++
		countStr := ""
		for i < length && input[i] >= '0' && input[i] <= '9' {
			countStr += string(input[i])
			i++
		}
		count, _ := strconv.Atoi(countStr)
		result.WriteString(strings.Repeat(string(char), count))
	}
	return result.String()
}

func CompressFile(inputFile, outputFile string) error {
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		content.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	compressed := Compress(content.String())

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.WriteString(compressed)
	if err != nil {
		return err
	}

	return nil
}

func DecompressFile(inputFile, outputFile string) error {
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		content.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	decompressed := Decompress(content.String())

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.WriteString(decompressed)
	if err != nil {
		return err
	}

	return nil
}
