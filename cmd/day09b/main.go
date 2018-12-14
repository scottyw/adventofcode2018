package main

import (
	"fmt"
)

type marble struct {
	value int
	left  *marble
	right *marble
}

func play(players, lastMarble int) int {
	var value = 1
	var current = &marble{value: 0}
	current.left = current
	current.right = current
	var scores = make([]int, players)
	for {
		if value > lastMarble {
			var max int
			for _, score := range scores {
				if score > max {
					max = score
				}
			}
			return max
		}
		if value%23 == 0 {
			for x := 0; x < 7; x++ {
				current = current.left
			}
			player := value % players
			scores[player] += value
			scores[player] += current.value
			// Delete current
			left := current.left
			right := current.right
			left.right = right
			right.left = left
			current = right
		} else {
			left := current.right
			right := current.right.right
			current = &marble{
				value: value,
				left:  left,
				right: right,
			}
			left.right = current
			right.left = current
		}
		value++
	}
}

func main() {
	fmt.Println(play(462, 7193800))
}
