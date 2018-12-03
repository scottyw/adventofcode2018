package main

import (
	"fmt"
	"regexp"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

const dim = 1200

func main() {
	lines := aoc.FileToLines("input/3.txt")
	// #22 @ 723,235: 22x21
	r := regexp.MustCompile("#\\d+ @ (\\d+),(\\d+): (\\d+)x(\\d+)")
	allocated := [dim][dim]bool{}
	clashes := [dim][dim]bool{}
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		ox := aoc.Atoi(matches[1])
		oy := aoc.Atoi(matches[2])
		wx := aoc.Atoi(matches[3])
		wy := aoc.Atoi(matches[4])
		for x := ox; x < ox+wx; x++ {
			for y := oy; y < oy+wy; y++ {
				isAllocated := allocated[x][y]
				if isAllocated {
					clashes[x][y] = true
				}
				allocated[x][y] = true
			}
		}
	}
	var total int
	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			if clashes[x][y] {
				total++
			}
		}
	}
	fmt.Println(total)
}
