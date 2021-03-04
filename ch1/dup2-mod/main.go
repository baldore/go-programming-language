// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type duplicatedLine struct {
	count     int
	fileNames []string
}

type countsMap = map[string]*duplicatedLine

func countLines(f *os.File, counts countsMap) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		_, ok := counts[key]

		if !ok {
			counts[key] = new(duplicatedLine)
		}

		counts[key].count++
		counts[key].fileNames = append(counts[key].fileNames, f.Name())
	}
}

func countLinesInFiles(files []string, counts countsMap) {
	for _, filename := range files {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}
}

func printCounts(counts countsMap) {
	for line, data := range counts {
		if data.count > 1 {
			fmt.Printf("%d\t%s\n", len(data.fileNames), data.fileNames)
			fmt.Printf("%d\t%s\n", data.count, line)
		}
	}
}

func main() {
	counts := make(countsMap)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		countLinesInFiles(files, counts)
	}

	printCounts(counts)
}
