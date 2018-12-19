package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

type train struct {
	x           int
	y           int
	orientation string
	lastTurn    int
}

var track = [150][150]rune{}

var trains = []*train{}

func init() {
	lines := aoc.FileToLines("input/13.txt")
	for y, line := range lines {
		for x, s := range line {
			switch s {
			case '<':
				track[x][y] = '-'
				trains = append(trains, &train{x: x, y: y, orientation: "w"})
			case '>':
				track[x][y] = '-'
				trains = append(trains, &train{x: x, y: y, orientation: "e"})
			case '^':
				track[x][y] = '|'
				trains = append(trains, &train{x: x, y: y, orientation: "n"})
			case 'v':
				track[x][y] = '|'
				trains = append(trains, &train{x: x, y: y, orientation: "s"})
			case '|', '-', '/', '\\', ' ', '+':
				track[x][y] = s
			default:
				panic(fmt.Sprintf("unknown track segment: %s", string(s)))
			}
		}
	}
}

func checkForCollision(train *train) *train {
	for _, candidate := range trains {
		if candidate == train {
			continue
		}
		if candidate.x == train.x && candidate.y == train.y {
			fmt.Printf("Collision! %p %v %p %v\n", train, train, candidate, candidate)
			return candidate
		}
	}
	return nil
}

func tick() {
	var collidingTrains = map[*train]bool{}
	for _, train := range trains {
		// Don't move trains that have already been in a collision
		if _, ok := collidingTrains[train]; ok {
			continue
		}
		// Move train based on track
		switch track[train.x][train.y] {
		case '|':
			if train.orientation == "n" {
				train.y--
			} else if train.orientation == "s" {
				train.y++
			} else {
				panic(fmt.Sprintf("unexpected orientation: %v", train))
			}
		case '-':
			if train.orientation == "w" {
				train.x--
			} else if train.orientation == "e" {
				train.x++
			} else {
				panic(fmt.Sprintf("unexpected orientation: %v", train))
			}
		case '\\':
			if train.orientation == "n" {
				train.orientation = "w"
				train.x--
			} else if train.orientation == "s" {
				train.orientation = "e"
				train.x++
			} else if train.orientation == "w" {
				train.orientation = "n"
				train.y--
			} else if train.orientation == "e" {
				train.orientation = "s"
				train.y++
			} else {
				panic(fmt.Sprintf("unexpected orientation: %v", train))
			}
		case '/':
			if train.orientation == "n" {
				train.orientation = "e"
				train.x++
			} else if train.orientation == "s" {
				train.orientation = "w"
				train.x--
			} else if train.orientation == "w" {
				train.orientation = "s"
				train.y++
			} else if train.orientation == "e" {
				train.orientation = "n"
				train.y--
			} else {
				panic(fmt.Sprintf("unexpected orientation: %v", train))
			}
		case '+':
			if train.orientation == "n" {
				switch train.lastTurn % 3 {
				case 0:
					train.orientation = "w"
					train.x--
				case 1:
					train.y--
				case 2:
					train.orientation = "e"
					train.x++
				}
			} else if train.orientation == "s" {
				switch train.lastTurn % 3 {
				case 0:
					train.orientation = "e"
					train.x++
				case 1:
					train.y++
				case 2:
					train.orientation = "w"
					train.x--
				}
			} else if train.orientation == "w" {
				switch train.lastTurn % 3 {
				case 0:
					train.orientation = "s"
					train.y++
				case 1:
					train.x--
				case 2:
					train.orientation = "n"
					train.y--
				}
			} else if train.orientation == "e" {
				switch train.lastTurn % 3 {
				case 0:
					train.orientation = "n"
					train.y--
				case 1:
					train.x++
				case 2:
					train.orientation = "s"
					train.y++
				}
			} else {
				panic(fmt.Sprintf("unexpected orientation: %v", train))
			}
			train.lastTurn++
		default:
			panic(fmt.Sprintf("unexpected track : %v", track[train.x][train.y]))
		}
		collidingTrain := checkForCollision(train)
		if collidingTrain != nil {
			collidingTrains[train] = true
			collidingTrains[collidingTrain] = true
		}
	}
	// Remove trains that were involved in a collision
	var survivingTrains = []*train{}
	for _, train := range trains {
		if _, ok := collidingTrains[train]; !ok {
			survivingTrains = append(survivingTrains, train)
		}
	}
	trains = survivingTrains
}

func main() {
	for {
		tick()
		if len(trains) == 1 {
			fmt.Println(trains[0])
			return
		}
	}
}
