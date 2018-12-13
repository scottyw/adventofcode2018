package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

var numbers []int

func init() {
	input := aoc.FileToString("input/8.txt")
	numberStrings := strings.Split(input, " ")
	numbers = make([]int, len(numberStrings))
	for i, numberString := range numberStrings {
		numbers[i] = aoc.Atoi(numberString)
	}
}

func parseNode() int {
	if len(numbers) < 3 {
		panic(fmt.Sprintf("Length too short: %d\n", len(numbers)))
	}
	childCount := numbers[0]
	metadataCount := numbers[1]
	numbers = numbers[2:]
	if childCount == 0 {
		var metadataTotal int
		for i := 0; i < metadataCount; i++ {
			metadataTotal += numbers[0]
			numbers = numbers[1:]
		}
		return metadataTotal
	}
	childrenTotals := make([]int, childCount)
	for i := 0; i < childCount; i++ {
		childrenTotals[i] = parseNode()
	}
	var metadataTotal int
	for i := 0; i < metadataCount; i++ {
		// Offset by 1 since our array is zero-based and the metadata indexing is one-based
		child := numbers[0] - 1
		if child >= 0 && child < len(childrenTotals) {
			metadataTotal += childrenTotals[child]
		}
		numbers = numbers[1:]
	}
	return metadataTotal
}

func main() {
	fmt.Println(parseNode())
}
