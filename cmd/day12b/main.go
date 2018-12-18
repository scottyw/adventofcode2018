package main

import "fmt"

// Treat each of the 5 character patterns as a binary number
//
// ..... => .
// ....# => .
// ...#. => #
// ...## => .
// ..#.. => #
// ..#.# => .
// ..##. => .
// ..### => #
// .#... => .
// .#..# => #
// .#.#. => .
// .#.## => .
// .##.. => #
// .##.# => #
// .###. => #
// .#### => #
// #.... => .
// #...# => .
// #..#. => #
// #..## => .
// #.#.. => #
// #.#.# => #
// #.##. => #
// #.### => #
// ##... => #
// ##..# => .
// ##.#. => #
// ##.## => #
// ###.. => #
// ###.# => .
// ####. => .
// ##### => .
//
// For each pattern, use that binary number as the index into an array that shows whether the chosen cell is switched on
//
var patterns = [32]bool{
	false, false, true, false, true, false, false, true,
	false, true, false, false, true, true, true, true,
	false, false, true, false, true, true, true, true,
	true, false, true, true, true, false, false, false}

const size = 1000010

var cells = [size * 2]bool{}

func init() {
	state := "#.##.###.#.##...##..#..##....#.#.#.#.##....##..#..####..###.####.##.#..#...#..######.#.....#..##...#"
	for i, cell := range state {
		if cell == 35 {
			cells[size+i] = true
		} else if cell != 46 {
			panic("Bad character in state")
		}
	}
}

func calculateCell(i int) bool {
	var pattern int
	if cells[i-2] {
		pattern += 16
	}
	if cells[i-1] {
		pattern += 8
	}
	if cells[i] {
		pattern += 4
	}
	if cells[i+1] {
		pattern += 2
	}
	if cells[i+2] {
		pattern++
	}
	return patterns[pattern]
}

func nextGeneration(generation int) {
	var next = [size * 2]bool{}
	// Living cells expand the previous generation by no more than 1 cell each direction
	for i := size - generation; i < size+100+generation; i++ {
		next[i] = calculateCell(i)
	}
	cells = next
}

func addLivingCells() int {
	var total int
	for i, cell := range cells {
		if cell {
			total += i - size
		}
	}
	return total
}

func main() {
	for generation := 1; generation <= 2000; generation++ {
		nextGeneration(generation)
		if generation > 100 && generation%100 == 0 {
			if generation*62+655 != addLivingCells() {
				panic("Wrong assumptions!")
			}
		}
	}
	fmt.Println(50000000000*62 + 655)
}
