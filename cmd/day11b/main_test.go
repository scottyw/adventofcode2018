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
	// For grid serial number 18, the largest total square (with a total power of 113) is 16x16 and has a top-left corner of 90,269, so its identifier is 90,269,16.
	x, y, size, power := maxCellPower(18)
	assert.Equal(t, 90, x)
	assert.Equal(t, 269, y)
	assert.Equal(t, 16, size)
	assert.Equal(t, 113, power)

	// For grid serial number 42, the largest total square (with a total power of 119) is 12x12 and has a top-left corner of 232,251, so its identifier is 232,251,12.
	x, y, size, power = maxCellPower(42)
	assert.Equal(t, 232, x)
	assert.Equal(t, 251, y)
	assert.Equal(t, 12, size)
	assert.Equal(t, 119, power)
}
