package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestRun(t *testing.T) {

	// If the Elves think their skill will improve after making 9 recipes, the scores of the ten recipes after the first nine on the scoreboard would be 5158916779 (highlighted in the last line of the diagram).
	assert.Equal(t, "5158916779", run(9))
	// After 5 recipes, the scores of the next ten would be 0124515891.
	assert.Equal(t, "0124515891", run(5))
	// After 18 recipes, the scores of the next ten would be 9251071085.
	assert.Equal(t, "9251071085", run(18))
	// After 2018 recipes, the scores of the next ten would be 5941429882.
	assert.Equal(t, "5941429882", run(2018))

}
