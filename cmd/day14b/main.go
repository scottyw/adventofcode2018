package main

import (
	"fmt"
	"math"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

type recipe struct {
	value int
	next  *recipe
}

var elf1, elf2, start, end *recipe

func reinit() {
	elf1 = &recipe{value: 3}
	elf2 = &recipe{value: 7}
	elf1.next = elf2
	elf2.next = elf1
	start = elf1
	end = elf2
}

func print() {
	current := start
	for {
		if current == elf1 {
			fmt.Printf("(%d)", current.value)
		} else if current == elf2 {
			fmt.Printf("[%d]", current.value)
		} else {
			fmt.Printf(" %d ", current.value)
		}
		current = current.next
		if current == start {
			break
		}
	}
	fmt.Println()
}

func tick() int {
	sum := elf1.value + elf2.value
	if sum >= 10 {
		new1 := &recipe{value: 1}
		new2 := &recipe{value: sum - 10}
		new1.next = new2
		new2.next = end.next
		end.next = new1
		end = new2
	} else {
		new := &recipe{value: sum}
		new.next = end.next
		end.next = new
		end = new
	}
	iterations := elf1.value + 1
	for x := 0; x < iterations; x++ {
		elf1 = elf1.next
	}
	iterations = elf2.value + 1
	for x := 0; x < iterations; x++ {
		elf2 = elf2.next
	}
	return sum
}

func run(searchString string) int {
	reinit()
	search := aoc.Atoi(searchString)
	denominator := int(math.Pow10(len(searchString)))
	last := 37
	count := 2
	current := start
	for {
		sum := tick()
		if sum >= 10 {
			// Check one digit at a time
			count++
			last *= 10
			last++ // The first digit of the sum, which must be 1
			last %= denominator
			if last == search {
				return count - len(searchString)
			}
			// Check second digit
			count++
			last *= 10
			last += sum % 10 // The second digit of the sum
			last %= denominator
			if last == search {
				return count - len(searchString)
			}
		} else {
			count++
			last *= 10
			last += sum
			last %= denominator
			if last == search {
				return count - len(searchString)
			}
		}
		current = current.next
	}
}

func main() {
	count := run("793061")
	fmt.Println(count)
}
