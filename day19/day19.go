package day19

import "github.com/AlbertBrand/aoc2017/util"

type Direction int

const (
	n Direction = iota
	e
	s
	w
)

type Grid struct {
	g [][]byte
}

func (g *Grid) get(x int, y int) (byte, bool) {
	if x < 0 || x > len(g.g[0])-1 ||
		y < 0 || y > len(g.g)-1 {
		return g.g[0][0], false
	}
	return g.g[y][x], true
}

func (g *Grid) newDirection(x int, y int, dir Direction) Direction {
	if dir == n || dir == s {
		if next, valid := g.get(x+1, y); valid && next != ' ' {
			return e
		}
		if next, valid := g.get(x-1, y); valid && next != ' ' {
			return w
		}
	} else {
		if next, valid := g.get(x, y-1); valid && next != ' ' {
			return n
		}
		if next, valid := g.get(x, y+1); valid && next != ' ' {
			return s
		}
	}
	return -1
}

func toGrid(lines []string) Grid {
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		gridline := make([]byte, len(line))
		grid[i] = gridline
		for j, char := range line {
			grid[i][j] = byte(char)
		}
	}
	return Grid{
		g: grid,
	}
}

func solver(lines []string) {
	grid := toGrid(lines)

	y := 0
	x := 0
	var char byte
	for x, char = range grid.g[y] {
		if char == '|' {
			break
		}
	}

	word := ""
	direction := s
	steps := 0

main:
	for {
		switch direction {
		case n:
			y -= 1
		case e:
			x += 1
		case s:
			y += 1
		case w:
			x -= 1
		}
		steps++

		elem, _ := grid.get(x, y)
		switch {
		case 'A' <= elem && elem <= 'Z':
			word += string(elem)
		case elem == '+':
			direction = grid.newDirection(x, y, direction)
		case elem == ' ':
			break main
		}
	}

	println(word)
	println(steps)
}

func Test() {
	solver(util.ReadTxt("day19/test.txt"))
}

func Solve() {
	solver(util.ReadTxt("day19/input.txt"))
}
