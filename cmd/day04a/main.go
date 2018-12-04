package main

import (
	"fmt"
	"regexp"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

const dim = 1200

func main() {
	// This implementation assumes the input is sorted, which I did using the editor
	lines := aoc.FileToLines("input/4.txt")
	// [1518-03-22 00:00] Guard #503 begins shift
	// [1518-03-22 00:17] falls asleep
	// [1518-03-22 00:36] wakes up
	// [1518-03-23 00:02] Guard #223 begins shift
	// [1518-03-23 00:23] falls asleep
	// [1518-03-23 00:55] wakes up
	guardRegex := regexp.MustCompile("\\[\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d\\] Guard #(\\d+)")
	sleepsRegex := regexp.MustCompile("\\[\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:(\\d\\d)\\] falls asleep")
	wakesRegex := regexp.MustCompile("\\[\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:(\\d\\d)\\] wakes up")
	var guard, sleeps, wakes int
	var matches []string
	sleepTracking := map[int]*[60]int{}
	for _, line := range lines {
		matches = guardRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			guard = aoc.Atoi(matches[1])
			continue
		}
		matches = sleepsRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			sleeps = aoc.Atoi(matches[1])
			continue
		}
		matches = wakesRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			wakes = aoc.Atoi(matches[1])
			shift, ok := sleepTracking[guard]
			if !ok {
				shift = &[60]int{}
				sleepTracking[guard] = shift
			}
			for min := sleeps; min < wakes; min++ {
				shift[min]++
			}
		}
	}
	var chosenGuard, chosenMinute, maxSleep int
	for guard, shift := range sleepTracking {
		var totalSleep, sleepiestMinuteTotal, sleepiestMinute int
		for min := 0; min < len(shift); min++ {
			totalSleep += shift[min]
			if shift[min] > sleepiestMinuteTotal {
				sleepiestMinuteTotal = shift[min]
				sleepiestMinute = min
			}
		}
		if totalSleep > maxSleep {
			chosenGuard = guard
			chosenMinute = sleepiestMinute
			maxSleep = totalSleep
		}
	}
	fmt.Println(chosenGuard)
	fmt.Println(chosenMinute)
	fmt.Println(chosenGuard * chosenMinute)
}
