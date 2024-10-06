// package huffman

// import (
// 	"bufio"
// 	"container/heap"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// // Node represents a node in the Huffman tree
// type Node struct {
// 	char  rune
// 	freq  int
// 	left  *Node
// 	right *Node
// }

// // PriorityQueue implements heap.Interface and holds Nodes
// type PriorityQueue []*Node

// func (pq PriorityQueue) Len() int           { return len(pq) }
// func (pq PriorityQueue) Less(i, j int) bool { return pq[i].freq < pq[j].freq }
// func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
// func (pq *PriorityQueue) Push(x interface{}) {
// 	*pq = append(*pq, x.(*Node))
// }
// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	x := old[n-1]
// 	*pq = old[0 : n-1]
// 	return x
// }

// // BuildHuffmanTree builds the Huffman tree from character frequencies
// func BuildHuffmanTree(freqs map[rune]int) *Node {
// 	pq := &PriorityQueue{}
// 	heap.Init(pq)

// 	for char, freq := range freqs {
// 		heap.Push(pq, &Node{char: char, freq: freq})
// 	}

// 	for pq.Len() > 1 {
// 		left := heap.Pop(pq).(*Node)
// 		right := heap.Pop(pq).(*Node)
// 		merged := &Node{
// 			char:  0,
// 			freq:  left.freq + right.freq,
// 			left:  left,
// 			right: right,
// 		}
// 		heap.Push(pq, merged)
// 	}

// 	return heap.Pop(pq).(*Node)
// }

// // BuildCodeTable generates a mapping from characters to their Huffman codes
// func BuildCodeTable(root *Node) map[rune]string {
// 	codes := make(map[rune]string)
// 	var encode func(*Node, string)
// 	encode = func(node *Node, code string) {
// 		if node == nil {
// 			return
// 		}
// 		if node.left == nil && node.right == nil {
// 			codes[node.char] = code
// 		}
// 		encode(node.left, code+"0")
// 		encode(node.right, code+"1")
// 	}
// 	encode(root, "")
// 	return codes
// }

// // Compress encodes the input string using the Huffman code table
// func Compress(input string, codes map[rune]string) string {
// 	var compressed strings.Builder
// 	for _, char := range input {
// 		compressed.WriteString(codes[char])
// 	}
// 	return compressed.String()
// }

// // Decompress decodes the compressed string using the Huffman tree
// func Decompress(compressed string, root *Node) string {
// 	var result strings.Builder
// 	node := root
// 	for _, bit := range compressed {
// 		if bit == '0' {
// 			node = node.left
// 		} else {
// 			node = node.right
// 		}
// 		if node.left == nil && node.right == nil {
// 			result.WriteRune(node.char)
// 			node = root
// 		}
// 	}
// 	return result.String()
// }

// func CompressFile(inputFile, outputFile string) error {
// 	inFile, err := os.Open(inputFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer inFile.Close()

// 	freqs := make(map[rune]int)
// 	scanner := bufio.NewScanner(inFile)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		for _, char := range line {
// 			freqs[char]++
// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	root := BuildHuffmanTree(freqs)
// 	codes := BuildCodeTable(root)

// 	inFile.Seek(0, io.SeekStart)
// 	var compressed strings.Builder
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		for _, char := range line {
// 			compressed.WriteString(codes[char])
// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	outFile, err := os.Create(outputFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer outFile.Close()

// 	writer := bufio.NewWriter(outFile)
// 	for char, code := range codes {
// 		fmt.Fprintf(writer, "%c:%s\n", char, code)
// 	}
// 	writer.WriteString("\n")
// 	writer.WriteString(compressed.String())
// 	writer.Flush()

// 	return nil
// }

// func DecompressFile(inputFile, outputFile string) error {
// 	inFile, err := os.Open(inputFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer inFile.Close()

// 	codes := make(map[string]rune)
// 	var root *Node
// 	scanner := bufio.NewScanner(inFile)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if line == "" {
// 			break
// 		}
// 		var char rune
// 		var code string
// 		fmt.Sscanf(line, "%c:%s", &char, &code)
// 		codes[code] = char
// 	}

// 	root = BuildHuffmanTreeFromCodes(codes)

// 	var compressed strings.Builder
// 	for scanner.Scan() {
// 		compressed.WriteString(scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	decompressed := Decompress(compressed.String(), root)

// 	outFile, err := os.Create(outputFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer outFile.Close()

// 	_, err = outFile.WriteString(decompressed)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // BuildHuffmanTreeFromCodes builds a Huffman tree from the code table
//
//	func BuildHuffmanTreeFromCodes(codes map[string]rune) *Node {
//		root := &Node{}
//		for code, char := range codes {
//			node := root
//			for _, bit := range code {
//				if bit == '0' {
//					if node.left == nil {
//						node.left = &Node{}
//					}
//					node = node.left
//				} else {
//					if node.right == nil {
//						node.right = &Node{}
//					}
//					node = node.right
//				}
//			}
//			node.char = char
//		}
//		return root
//	}
// package huffman

// import (
// 	"bufio"
// 	"container/heap"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// // Node represents a node in the Huffman tree
// type Node struct {
// 	char  rune
// 	freq  int
// 	left  *Node
// 	right *Node
// }

// // PriorityQueue implements heap.Interface and holds Nodes
// type PriorityQueue []*Node

// func (pq PriorityQueue) Len() int           { return len(pq) }
// func (pq PriorityQueue) Less(i, j int) bool { return pq[i].freq < pq[j].freq }
// func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
// func (pq *PriorityQueue) Push(x interface{}) {
// 	*pq = append(*pq, x.(*Node))
// }
// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	x := old[n-1]
// 	*pq = old[0 : n-1]
// 	return x
// }

// // BuildHuffmanTree builds the Huffman tree from character frequencies
// func BuildHuffmanTree(freqs map[rune]int) *Node {
// 	pq := &PriorityQueue{}
// 	heap.Init(pq)

// 	for char, freq := range freqs {
// 		heap.Push(pq, &Node{char: char, freq: freq})
// 	}

// 	for pq.Len() > 1 {
// 		left := heap.Pop(pq).(*Node)
// 		right := heap.Pop(pq).(*Node)
// 		merged := &Node{
// 			char:  0,
// 			freq:  left.freq + right.freq,
// 			left:  left,
// 			right: right,
// 		}
// 		heap.Push(pq, merged)
// 	}

// 	return heap.Pop(pq).(*Node)
// }

// // BuildCodeTable generates a mapping from characters to their Huffman codes
// func BuildCodeTable(root *Node) map[rune]string {
// 	codes := make(map[rune]string)
// 	var encode func(*Node, string)
// 	encode = func(node *Node, code string) {
// 		if node == nil {
// 			return
// 		}
// 		if node.left == nil && node.right == nil {
// 			codes[node.char] = code
// 		}
// 		encode(node.left, code+"0")
// 		encode(node.right, code+"1")
// 	}
// 	encode(root, "")
// 	return codes
// }

// // Compress encodes the input string using the Huffman code table
// func Compress(input string, codes map[rune]string) string {
// 	var compressed strings.Builder
// 	for _, char := range input {
// 		compressed.WriteString(codes[char])
// 	}
// 	return compressed.String()
// }

// // Decompress decodes the compressed string using the Huffman tree
// func Decompress(compressed string, root *Node) string {
// 	var result strings.Builder
// 	node := root
// 	for _, bit := range compressed {
// 		if bit == '0' {
// 			node = node.left
// 		} else {
// 			node = node.right
// 		}
// 		if node.left == nil && node.right == nil {
// 			result.WriteRune(node.char)
// 			node = root
// 		}
// 	}
// 	return result.String()
// }

// func CompressFile(inputFile, outputFile string) error {
// 	inFile, err := os.Open(inputFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer inFile.Close()

// 	freqs := make(map[rune]int)
// 	reader := bufio.NewReader(inFile)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err != nil && err != io.EOF {
// 			return err
// 		}
// 		for _, char := range line {
// 			freqs[char]++
// 		}
// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	root := BuildHuffmanTree(freqs)
// 	codes := BuildCodeTable(root)

// 	// Reset file pointer to read the file again for compression
// 	inFile.Seek(0, io.SeekStart)

// 	var compressed strings.Builder
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err != nil && err != io.EOF {
// 			return err
// 		}
// 		for _, char := range line {
// 			compressed.WriteString(codes[char])
// 		}
// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	outFile, err := os.Create(outputFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer outFile.Close()

// 	writer := bufio.NewWriter(outFile)
// 	for char, code := range codes {
// 		fmt.Fprintf(writer, "%c:%s\n", char, code)
// 	}
// 	writer.WriteString("\n")
// 	writer.WriteString(compressed.String())
// 	writer.Flush()

// 	return nil
// }

// func DecompressFile(inputFile, outputFile string) error {
// 	inFile, err := os.Open(inputFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer inFile.Close()

// 	codes := make(map[string]rune)
// 	var root *Node
// 	reader := bufio.NewReader(inFile)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err != nil && err != io.EOF {
// 			return err
// 		}
// 		if line == "\n" {
// 			break
// 		}
// 		var char rune
// 		var code string
// 		fmt.Sscanf(line, "%c:%s", &char, &code)
// 		codes[code] = char
// 	}

// 	root = BuildHuffmanTreeFromCodes(codes)

// 	var compressed strings.Builder
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err != nil && err != io.EOF {
// 			return err
// 		}
// 		compressed.WriteString(line)
// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	decompressed := Decompress(compressed.String(), root)

// 	outFile, err := os.Create(outputFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer outFile.Close()

// 	_, err = outFile.WriteString(decompressed)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // BuildHuffmanTreeFromCodes builds a Huffman tree from the code table
// func BuildHuffmanTreeFromCodes(codes map[string]rune) *Node {
// 	root := &Node{}
// 	for code, char := range codes {
// 		node := root
// 		for _, bit := range code {
// 			if bit == '0' {
// 				if node.left == nil {
// 					node.left = &Node{}
// 				}
// 				node = node.left
// 			} else {
// 				if node.right == nil {
// 					node.right = &Node{}
// 				}
// 				node = node.right
// 			}
// 		}
// 		node.char = char
// 	}
// 	return root
// }

package huffman

import (
	"container/heap"
	"encoding/gob"
	"fmt"
	"os"
)

type Node struct {
	Char  byte
	Freq  int
	Left  *Node
	Right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Freq < pq[j].Freq }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type BitWriter struct {
	file    *os.File
	buffer  byte
	counter uint8
}

func NewBitWriter(file *os.File) *BitWriter {
	return &BitWriter{file: file}
}

func (w *BitWriter) WriteBit(bit bool) error {
	if bit {
		w.buffer |= 1 << (7 - w.counter)
	}
	w.counter++
	if w.counter == 8 {
		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}

func (w *BitWriter) Flush() error {
	if w.counter > 0 {
		if err := w.WriteByte(w.buffer); err != nil {
			return err
		}
		w.buffer = 0
		w.counter = 0
	}
	return nil
}

func (w *BitWriter) WriteByte(b byte) error {
	_, err := w.file.Write([]byte{b})
	return err
}

type BitReader struct {
	file    *os.File
	buffer  byte
	counter uint8
}

func NewBitReader(file *os.File) *BitReader {
	return &BitReader{file: file}
}

func (r *BitReader) ReadBit() (bool, error) {
	if r.counter == 0 {
		buffer := make([]byte, 1)
		_, err := r.file.Read(buffer)
		if err != nil {
			return false, err
		}
		r.buffer = buffer[0]
		r.counter = 8
	}
	r.counter--
	return (r.buffer & (1 << r.counter)) != 0, nil
}

func CompressFile(inputPath, outputPath string) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	freqTable := buildFrequencyTable(input)
	root := buildHuffmanTree(freqTable)
	codes := make(map[byte]string)
	buildHuffmanCodes(root, "", codes)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	// Write the tree structure for decompression
	encoder := gob.NewEncoder(outputFile)
	if err := encoder.Encode(root); err != nil {
		return fmt.Errorf("failed to encode tree: %v", err)
	}

	// Write the compressed data
	writer := NewBitWriter(outputFile)
	for _, b := range input {
		code := codes[b]
		for _, bit := range code {
			if err := writer.WriteBit(bit == '1'); err != nil {
				return fmt.Errorf("failed to write bit: %v", err)
			}
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %v", err)
	}

	return nil
}

func DecompressFile(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer inputFile.Close()

	// Read the tree structure
	decoder := gob.NewDecoder(inputFile)
	var root Node
	if err := decoder.Decode(&root); err != nil {
		return fmt.Errorf("failed to decode tree: %v", err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	reader := NewBitReader(inputFile)
	current := &root
	for {
		bit, err := reader.ReadBit()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("failed to read bit: %v", err)
		}

		if bit {
			current = current.Right
		} else {
			current = current.Left
		}

		if current.Left == nil && current.Right == nil {
			if _, err := outputFile.Write([]byte{current.Char}); err != nil {
				return fmt.Errorf("failed to write character: %v", err)
			}
			current = &root
		}
	}

	return nil
}

// Helper functions remain the same as in your original code
func buildFrequencyTable(data []byte) map[byte]int {
	freqTable := make(map[byte]int)
	for _, b := range data {
		freqTable[b]++
	}
	return freqTable
}

func buildHuffmanTree(freqTable map[byte]int) *Node {
	pq := make(PriorityQueue, 0)
	for char, freq := range freqTable {
		pq = append(pq, &Node{Char: char, Freq: freq})
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)
		merged := &Node{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}
		heap.Push(&pq, merged)
	}

	return heap.Pop(&pq).(*Node)
}

func buildHuffmanCodes(node *Node, code string, codes map[byte]string) {
	if node.Left == nil && node.Right == nil {
		codes[node.Char] = code
		return
	}
	if node.Left != nil {
		buildHuffmanCodes(node.Left, code+"0", codes)
	}
	if node.Right != nil {
		buildHuffmanCodes(node.Right, code+"1", codes)
	}
}
