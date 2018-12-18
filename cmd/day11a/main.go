package main

import "fmt"

func power(serialNumber, x, y int) int {
	// Find the fuel cell's rack ID, which is its X coordinate plus 10.
	rackID := x + 10
	// Begin with a power level of the rack ID times the Y coordinate.
	power := rackID * y
	// Increase the power level by the value of the grid serial number (your puzzle input).
	power += serialNumber
	// Set the power level to itself multiplied by the rack ID.
	power *= rackID
	// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
	power = (power / 100) % 10
	// Subtract 5 from the power level.
	power -= 5
	return power
}

func cellPower(serialNumber, x, y int) int {
	var total int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			total += power(serialNumber, x+i, y+j)
		}
	}
	return total
}

func maxCellPower(serialNumber int) (x, y, power int) {
	for i := 1; i <= 298; i++ {
		for j := 1; j <= 298; j++ {
			cellPower := cellPower(serialNumber, i, j)
			if cellPower > power {
				power = cellPower
				x = i
				y = j
			}
		}
	}
	return
}

func main() {
	x, y, power := maxCellPower(3999)
	fmt.Printf("Max power was %d found at cell %d,%d\n", power, x, y)
}
