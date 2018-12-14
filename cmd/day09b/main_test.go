package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestPlay(t *testing.T) {
	// 	10 players; last marble is worth 1618 points: high score is 8317
	assert.Equal(t, 8317, play(10, 1618))
	// 	13 players; last marble is worth 7999 points: high score is 146373
	assert.Equal(t, 146373, play(13, 7999))
	// 	17 players; last marble is worth 1104 points: high score is 2764
	assert.Equal(t, 2764, play(17, 1104))
	// 	21 players; last marble is worth 6111 points: high score is 54718
	assert.Equal(t, 54718, play(21, 6111))
	// 	30 players; last marble is worth 5807 points: high score is 37305
	assert.Equal(t, 37305, play(30, 5807))
}
