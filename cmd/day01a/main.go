package main

import (
	"fmt"
	"strconv"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

func main() {
	total := 0
	lines := aoc.FileToLines("input/1.txt")
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		aoc.Check(err)
		total += i
	}
	fmt.Println(total)
}
