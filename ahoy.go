// Ahoy is a tool to automate the generation of code for basic CRUD operations
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	output = flag.String("output", "", "output directory for the code; by default the current directory")
	input  = flag.String("input", "", "the file to read ")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tahoy [flags]\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("ahoy: ")
	flag.Usage = Usage
	flag.Parse()
	decoder, err := InitDecoder(*input)
	if *input == "" {
		flag.Usage()
		os.Exit(2)
	}
	if err != nil {
		log.Fatalf("Unrecognized file, ERROR: %v", err)
	}

	inputFile, err := os.Open(*input)
	if err != nil {
		log.Fatalf("Error reading file %v, ERROR: %v", *input, err)
	}
	defer inputFile.Close()
	project := &Project{}
	project.load(inputFile, decoder)
}
