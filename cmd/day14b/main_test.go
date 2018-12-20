package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestRun(t *testing.T) {

	// 51589 first appears after 9 recipes.
	assert.Equal(t, 9, run("515891"))
	assert.Equal(t, 9, run("51589"))
	assert.Equal(t, 9, run("5158"))
	assert.Equal(t, 9, run("515"))
	// 01245 first appears after 5 recipes.
	assert.Equal(t, 5, run("01245"))
	// 92510 first appears after 18 recipes.
	assert.Equal(t, 18, run("92510"))
	// 59414 first appears after 2018 recipes.
	assert.Equal(t, 2018, run("59414"))

}
