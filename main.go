package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go input.csv output.csv")
		return
	}
	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()
	reader := csv.NewReader(inputFile)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading input file as CSV:", err)
		return
	}
	header := lines[0]
	lines = lines[1:]
	sort.SliceStable(lines, func(i, j int) bool {
		for col := 0; col < len(header); col++ {
			if lines[i][col] != lines[j][col] {
				return lines[i][col] < lines[j][col]
			}
		}
		return false
	})
	lines = append([][]string{header}, lines...)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	err = writer.WriteAll(lines)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
}
