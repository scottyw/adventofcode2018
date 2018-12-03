package aoc

import (
	"bufio"
	"os"
	"strconv"
)

// Check panics if the error is not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Atoi converts a string to an int and panics if it goes wrong
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
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
