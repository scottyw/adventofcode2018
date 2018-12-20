package main

import "fmt"

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

func tick() {
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
}

func run(count int) string {
	reinit()
	current := start
	digits := ""
	for x := 0; x < count+10; x++ {
		tick()
		// print()
		if x >= count {
			digits += fmt.Sprintf("%d", current.value)
		}
		current = current.next
	}
	return digits
}

func main() {
	digits := run(793061)
	fmt.Println(digits)
}
