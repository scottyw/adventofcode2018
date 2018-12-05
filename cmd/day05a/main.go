package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

func convertToInt(b byte) int {
	i := int(b)
	if i < 91 {
		return -i + 64
	}
	return i - 96
}

func convertToIntArray(bs []byte) []int {
	is := make([]int, len(bs))
	for index, b := range bs {
		is[index] = convertToInt(b)
	}
	return is
}

func destroyFirstPair(is []int) []int {
	if len(is) < 2 {
		return is
	}
	for index := 0; index < len(is)-1; index++ {
		if is[index]+is[index+1] == 0 {
			return append(is[:index], is[index+2:]...)
		}
	}
	return is
}

func main() {
	bs := aoc.FileToBytes("input/5.txt")
	is := convertToIntArray(bs)
	for {
		updated := destroyFirstPair(is)
		if len(updated) == len(is) {
			fmt.Println(len(updated))
			return
		}
		is = updated
	}
}
