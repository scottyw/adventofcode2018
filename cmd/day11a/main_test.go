package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestPower(t *testing.T) {
	assert.Equal(t, 4, power(8, 3, 5))

	// Fuel cell at  122,79, grid serial number 57: power level -5.
	assert.Equal(t, -5, power(57, 122, 79))

	// Fuel cell at 217,196, grid serial number 39: power level  0.
	assert.Equal(t, 0, power(39, 217, 196))

	// Fuel cell at 101,153, grid serial number 71: power level  4.
	assert.Equal(t, 4, power(71, 101, 153))
}

func TestMaxCellPower(t *testing.T) {
	// For grid serial number 18, the largest total 3x3 square has a top-left corner of 33,45 (with a total power of 29)
	x, y, power := maxCellPower(18)
	assert.Equal(t, 29, power)
	assert.Equal(t, 33, x)
	assert.Equal(t, 45, y)

	// For grid serial number 42, the largest 3x3 square's top-left is 21,61 (with a total power of 30)
	x, y, power = maxCellPower(42)
	assert.Equal(t, 30, power)
	assert.Equal(t, 21, x)
	assert.Equal(t, 61, y)
}
