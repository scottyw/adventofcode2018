package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

func main() {
	var exactly2, exactly3 int
	lines := aoc.FileToLines("input/2.txt")
	for _, line := range lines {
		counts := map[rune]int{}
		for _, c := range line {
			count, ok := counts[c]
			if ok {
				count++
				counts[c] = count
			} else {
				counts[c] = 1
			}
		}
		var found2, found3 bool
		for _, count := range counts {
			if count == 2 {
				found2 = true
			}
			if count == 3 {
				found3 = true
			}
		}
		if found2 {
			exactly2++
		}
		if found3 {
			exactly3++
		}
	}
	fmt.Println(exactly2 * exactly3)
}
