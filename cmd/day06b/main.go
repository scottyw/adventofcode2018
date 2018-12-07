package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

func calulcateDistance(i, j, x, y int) int {
	distance := 0
	if i > x {
		distance += i - x
	} else {
		distance += x - i
	}
	if j > y {
		distance += j - y
	} else {
		distance += y - j
	}
	return distance
}

func main() {
	lines := aoc.FileToLines("input/6.txt")
	// Assume the region we're looking for is contiguous and that we're not looking for the biggest of several
	count := 0
	// The coords are in the range 0-400, let's assume that's still fine for part B
	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			totalDistance := 0
			for _, line := range lines {
				// This is expensive but not so expensive it matters ...
				coords := strings.Split(line, ", ")
				x := aoc.Atoi(coords[0])
				y := aoc.Atoi(coords[1])
				totalDistance += calulcateDistance(i, j, x, y)
			}
			if totalDistance < 10000 {
				count++
			}
		}
	}
	fmt.Println(count)
}
