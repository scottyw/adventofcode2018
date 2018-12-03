package main

import (
	"fmt"
	"strconv"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

func main() {
	total := 0
	seen := map[int]bool{}
	lines := aoc.FileToLines("input/1.txt")
	for {
		for _, line := range lines {
			if _, ok := seen[total]; ok {
				fmt.Println(total)
				return
			} else {
				seen[total] = true
			}
			i, err := strconv.Atoi(line)
			aoc.Check(err)
			total += i
		}
	}
}
