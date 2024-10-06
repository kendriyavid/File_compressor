package main

import (
	"fcompressor/cmd" // Replace with your actual module name if different
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
