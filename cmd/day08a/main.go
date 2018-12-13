package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

var numbers []int
var metadataTotal int

func init() {
	input := aoc.FileToString("input/8.txt")
	numberStrings := strings.Split(input, " ")
	numbers = make([]int, len(numberStrings))
	for i, numberString := range numberStrings {
		numbers[i] = aoc.Atoi(numberString)
	}
}

func parseNode() {
	if len(numbers) < 3 {
		panic(fmt.Sprintf("Length too short: %d\n", len(numbers)))
	}
	childCount := numbers[0]
	metadataCount := numbers[1]
	numbers = numbers[2:]
	for i := 0; i < childCount; i++ {
		parseNode()
	}
	for i := 0; i < metadataCount; i++ {
		metadataTotal += numbers[0]
		numbers = numbers[1:]
	}
}

func main() {
	parseNode()
	fmt.Println(metadataTotal)
}
