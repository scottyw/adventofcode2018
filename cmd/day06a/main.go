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
	area := [400][400]int{}
	// There are 50 inputs
	counts := [50]int{}
	// The coords are in the range 0-400
	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			closestDistance := 400
			closestIndex := 0
			for index, line := range lines {
				// This is expensive but not so expensive it matters ...
				coords := strings.Split(line, ", ")
				x := aoc.Atoi(coords[0])
				y := aoc.Atoi(coords[1])
				distance := calulcateDistance(i, j, x, y)
				if distance == closestDistance {
					// -1 means that more than one coord are equally close to this square
					closestIndex = -1
				}
				if distance < closestDistance {
					closestIndex = index
					closestDistance = distance
				}
			}
			if closestIndex != -1 {
				area[i][j] = closestIndex
				counts[closestIndex]++
			}
		}
	}
	// Any indexes at the edge of the area are infinite in size
	for x := 0; x < 400; x++ {
		counts[area[x][0]] = 0
		counts[area[x][399]] = 0
		counts[area[0][x]] = 0
		counts[area[399][x]] = 0
	}
	// Print them out and look for the largest ...
	for _, count := range counts {
		fmt.Println(count)
	}
}
