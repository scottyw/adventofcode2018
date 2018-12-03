package aoc

import (
	"bufio"
	"os"
)

// Check panics if the error is not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// FileToLines reads a file into an array of strings
func FileToLines(filePath string) []string {
	f, err := os.Open(filePath)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())
	return lines
}
