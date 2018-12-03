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
	r := regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")
	allocated := [dim][dim]int{}
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		id := aoc.Atoi(matches[1])
		ox := aoc.Atoi(matches[2])
		oy := aoc.Atoi(matches[3])
		wx := aoc.Atoi(matches[4])
		wy := aoc.Atoi(matches[5])
		for x := ox; x < ox+wx; x++ {
			for y := oy; y < oy+wy; y++ {
				allocated[x][y] = id
			}
		}
	}
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		id := aoc.Atoi(matches[1])
		ox := aoc.Atoi(matches[2])
		oy := aoc.Atoi(matches[3])
		wx := aoc.Atoi(matches[4])
		wy := aoc.Atoi(matches[5])
		var clash bool
		for x := ox; x < ox+wx; x++ {
			for y := oy; y < oy+wy; y++ {
				if allocated[x][y] != 0 && allocated[x][y] != id {
					clash = true
					allocated[x][y] = id
				}
			}
		}
		if !clash {
			fmt.Println(line)
		}
	}
}
