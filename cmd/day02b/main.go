package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

func main() {
	lines := aoc.FileToLines("input/2.txt")
	for _, line1 := range lines {
		for _, line2 := range lines {
			var diffs int
			for i := 0; i < len(line1); i++ {
				if line1[i] != line2[i] {
					diffs++
				}
			}
			if diffs == 1 {
				fmt.Println(line1)
				fmt.Println(line2)
				fmt.Println()
			}
		}
	}
}
