package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inputFile := flag.String("input", "", "Input file to read from (optional)")
	outputFile := flag.String("output", "", "Output file to write to (optional)")
	flag.Parse()

	var inputText string

	// Read from file or stdin
	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		buf := new(strings.Builder)
		_, err = io.Copy(buf, file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
			os.Exit(1)
		}
		inputText = buf.String()
	} else {
		reader := bufio.NewReader(os.Stdin)
		buf := new(strings.Builder)
		for {
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
				os.Exit(1)
			}
			buf.WriteString(line)
			if err == io.EOF {
				break
			}
		}
		inputText = buf.String()
	}

	// Compress the input text
	compressed := compressedString(inputText)

	// Write to file or stdout
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		_, err = file.WriteString(compressed)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to output file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(compressed)
	}
}

func compressedString(word string) string {
	var output strings.Builder
	compress(word, &output)
	return output.String()
}

func compress(word string, output *strings.Builder) {
	if len(word) == 0 {
		return
	}

	selected := word[0]
	counter := 1

	for ; counter < len(word); counter++ {
		if word[counter] != selected || counter > 8 {
			break
		}
	}

	output.WriteString(fmt.Sprintf("%d%c", counter, selected))

	compress(word[counter:], output)
}
