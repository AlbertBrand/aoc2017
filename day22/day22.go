package day22

import (
	"github.com/AlbertBrand/aoc2017/util"
)

type Coord struct {
	x, y int
}

type Direction uint8

const (
	up Direction = iota
	right
	down
	left
)

func solver(lines []string) {
	infected := make(map[Coord]bool)
	middle := (len(lines[0]) - 1) / 2
	for j, line := range lines {
		for i, node := range line {
			if node == '#' {
				c := Coord{i - middle, j - middle}
				infected[c] = true
			}
		}
	}

	direction := up
	pos := Coord{0, 0}
	count := 0
	for i := 0; i < 10000; i++ {
		if infected[pos] {
			direction = (direction + 1) % 4
			infected[pos] = false
		} else {
			direction = (direction - 1) % 4
			infected[pos] = true
			count++
		}
		switch direction {
		case up:
			pos = Coord{pos.x, pos.y - 1}
		case right:
			pos = Coord{pos.x + 1, pos.y}
		case down:
			pos = Coord{pos.x, pos.y + 1}
		case left:
			pos = Coord{pos.x - 1, pos.y}
		}
	}
	println(count)
}

type State uint8

const (
	C State = iota
	W
	I
	F
)

func solver2(lines []string) {
	infected := make(map[Coord]State)
	middle := (len(lines[0]) - 1) / 2
	for j, line := range lines {
		for i, node := range line {
			if node == '#' {
				c := Coord{i - middle, j - middle}
				infected[c] = I
			}
		}
	}

	direction := up
	pos := Coord{0, 0}
	count := 0
	for i := 0; i < 10000000; i++ {
		switch infected[pos] {
		case C:
			direction = (direction - 1) % 4
			infected[pos] = W
		case W:
			infected[pos] = I
			count++
		case I:
			direction = (direction + 1) % 4
			infected[pos] = F
		case F:
			direction = (direction + 2) % 4
			infected[pos] = C
		}
		switch direction {
		case up:
			pos = Coord{pos.x, pos.y - 1}
		case right:
			pos = Coord{pos.x + 1, pos.y}
		case down:
			pos = Coord{pos.x, pos.y + 1}
		case left:
			pos = Coord{pos.x - 1, pos.y}
		}
	}
	println(count)
}

func TestFirst() {
	solver(util.ReadTxt("day22/test.txt"))
}

func SolveFirst() {
	solver(util.ReadTxt("day22/input.txt"))
}

func TestSecond() {
	solver2(util.ReadTxt("day22/test.txt"))
}

func SolveSecond() {
	solver2(util.ReadTxt("day22/input.txt"))
}
