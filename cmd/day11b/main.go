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

func cellPower(cells *[301][301]int, serialNumber, x, y, size int) int {
	var total int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			total += cells[x+i][y+j]
		}
	}
	return total
}

func maxCellPower(serialNumber int) (x, y, size, p int) {
	cells := [301][301]int{}
	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			cells[i][j] = power(serialNumber, i, j)
		}
	}
	for s := 1; s <= 300; s++ {
		for i := 1; i <= 301-s; i++ {
			for j := 1; j <= 301-s; j++ {
				cellPower := cellPower(&cells, serialNumber, i, j, s)
				if cellPower > p {
					p = cellPower
					x = i
					y = j
					size = s
				}
			}
		}
	}
	return
}

func main() {
	x, y, size, power := maxCellPower(3999)
	fmt.Printf("Max power was %d found at cell %d,%d,%d\n", power, x, y, size)
}
