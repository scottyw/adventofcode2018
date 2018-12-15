package main

import (
	"fmt"
	"regexp"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

const dim = 500

var x []int
var y []int
var dx []int
var dy []int

func init() {
	lines := aoc.FileToLines("input/10.txt")
	// position=<10, -3> velocity=<-1,  1>
	r := regexp.MustCompile("position=<\\s*(-?\\d+),\\s*(-?\\d+)> velocity=<\\s*(-?\\d+),\\s*(-?\\d+)>")
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		x = append(x, aoc.Atoi(matches[1]))
		y = append(y, aoc.Atoi(matches[2]))
		dx = append(dx, aoc.Atoi(matches[3]))
		dy = append(dy, aoc.Atoi(matches[4]))
	}
}

func render() {
	var sky [dim][dim]bool
	var visible bool
	for i := 0; i < len(x); i++ {
		skyX := x[i]
		skyY := y[i]
		if skyX >= 0 && skyX < dim && skyY >= 0 && skyY < dim {
			sky[skyX][skyY] = true
			visible = true
		}
	}
	if visible {
		fmt.Println()
		// Narrow down these ranges by hand!
		for y := 120; y < 135; y++ {
			for x := 150; x < 250; x++ {
				if sky[x][y] {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
}

func update() {
	for i := 0; i < len(x); i++ {
		x[i] += dx[i]
		y[i] += dy[i]
	}
}

func main() {
	// Narrow down these ranges by hand!
	for x := 0; x < 12000; x++ {
		if x > 10344 && x < 10346 {
			fmt.Println("ROUND", x)
			render()
		}
		update()
	}
}
