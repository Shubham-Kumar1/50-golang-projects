package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/vjeantet/grok"
)

func main() {
	// Read pattern from pattern.txt
	patternBytes, err := os.ReadFile("pattern.txt")
	if err != nil {
		fmt.Printf("Error reading pattern file: %v\n", err)
		return
	}
	pattern := strings.TrimSpace(string(patternBytes))

	// Initialize grok
	g, err := grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
	if err != nil {
		fmt.Printf("Error initializing grok: %v\n", err)
		return
	}

	// Open input log file
	inputFile, err := os.Open("text.log")
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inputFile.Close()

	// Create output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Create a scanner to read the log file line by line
	scanner := bufio.NewScanner(inputFile)

	// Process each line
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines
		if line == "" {
			continue
		}

		// Parse the line using grok pattern
		values, err := g.Parse(pattern, line)
		if err != nil {
			fmt.Printf("Error parsing line %d: %v\n", lineNum, err)
			continue
		}

		// Convert parsed values to JSON
		jsonData, err := json.MarshalIndent(values, "", "  ")
		if err != nil {
			fmt.Printf("Error converting to JSON at line %d: %v\n", lineNum, err)
			continue
		}

		// Write to output file
		_, err = outputFile.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Printf("Error writing to output file at line %d: %v\n", lineNum, err)
			continue
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	fmt.Println("Log parsing completed successfully!")
}
