package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

func play(players, lastMarble int) int {
	var marble = 6
	var circle = []int{0, 4, 2, 5, 1, 3} // [5]  0  4  2 (5) 1  3
	var current = 3
	var scores = make([]int, players)
	for {
		if marble > lastMarble {
			var max int
			for _, score := range scores {
				if score > max {
					max = score
				}
			}
			return max
		}
		if marble%23 == 0 {
			player := marble % players
			scores[player] += marble
			current = (current - 7) % len(circle)
			if current < 0 {
				current += len(circle)
			}
			scores[player] += circle[current]
			circle = aoc.Delete(circle, current)
		} else {
			current = (current + 2) % len(circle)
			if current == 0 {
				current = len(circle)
			}
			circle = aoc.Insert(circle, current, marble)
		}
		marble++
	}
}

func main() {
	fmt.Println(play(462, 71938))
}
