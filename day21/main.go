package day21

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AlbertBrand/aoc2017/util"
)

var re = regexp.MustCompile(`([.#/]+) => ([.#/]+)`)

const start = ".#./..#/###"

// 0 1    1 3
// 2 3 => 0 2
func rotate2(s string) string {
	s0, s1, s2, s3 := s[0], s[1], s[3], s[4]
	return fmt.Sprintf("%c%c/%c%c", s1, s3, s0, s2)
}

// 0 1 2    2 5 8
// 3 4 5 => 1 4 7
// 6 7 8    0 3 6
func rotate3(s string) string {
	s0, s1, s2, s3, s4, s5, s6, s7, s8 := s[0], s[1], s[2], s[4], s[5], s[6], s[8], s[9], s[10]
	return fmt.Sprintf("%c%c%c/%c%c%c/%c%c%c", s2, s5, s8, s1, s4, s7, s0, s3, s6)
}

// 0 1 2    2 1 0
// 3 4 5 => 5 4 3
// 6 7 8    8 7 6
func fliph3(s string) string {
	s0, s1, s2, s3, s4, s5, s6, s7, s8 := s[0], s[1], s[2], s[4], s[5], s[6], s[8], s[9], s[10]
	return fmt.Sprintf("%c%c%c/%c%c%c/%c%c%c", s2, s1, s0, s5, s4, s3, s8, s7, s6)
}

func variations(s string) []string {
	v := make([]string, 0)
	v = append(v, s)

	if strings.Count(s, "/") == 1 {
		s = rotate2(s)
		v = append(v, s)
		s = rotate2(s)
		v = append(v, s)
		s = rotate2(s)
		v = append(v, s)
	} else {
		v = append(v, fliph3(s))
		s = rotate3(s)
		v = append(v, s)
		v = append(v, fliph3(s))
		s = rotate3(s)
		v = append(v, s)
		v = append(v, fliph3(s))
		s = rotate3(s)
		v = append(v, s)
		v = append(v, fliph3(s))
	}
	return v
}

type Grid [][]byte

func makeGrid(size int) Grid {
	grid := make(Grid, size)
	for a := range grid {
		grid[a] = make([]byte, size)
	}
	return grid
}

func draw(g Grid, s string, x int, y int) {
	lines := strings.Split(s, "/")
	for j, line := range lines {
		for i, b := range []byte(line) {
			g[y+j][x+i] = b
		}
	}
}

func getKey(g Grid, step int, x int, y int) string {
	s := ""
	for j := 0; j < step; j++ {
		for i := 0; i < step; i++ {
			s += string(g[y+j][x+i])
		}
		s += "/"
	}
	return s[:len(s)-1]
}

func solver(lines []string, iterations int) {
	var rules = make(map[string]string)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		from := match[1]
		to := match[2]
		for _, key := range variations(from) {
			rules[key] = to
		}
	}

	grid := makeGrid(3)
	draw(grid, start, 0, 0)

	for i := 0; i < iterations; i++ {
		size := len(grid[0])
		var newGrid Grid
		var step, stepNew int
		if size%2 == 0 {
			step = 2
			stepNew = 3
		} else {
			step = 3
			stepNew = 4
		}
		newGrid = makeGrid(size / step * stepNew)
		for y := 0; y < size; y += step {
			for x := 0; x < size; x += step {
				key := getKey(grid, step, x, y)
				if next, exists := rules[key]; !exists {
					panic("rule for " + key + " not found")
				} else {
					draw(newGrid, next, x/step*stepNew, y/step*stepNew)
				}
			}
		}
		grid = newGrid
	}
	count := 0
	for y, rows := range grid {
		for x := range rows {
			if grid[y][x] == '#' {
				count++
			}
		}
	}
	println(count)
}

func TestFirst() {
	solver([]string{"../.# => ##./#../...", ".#./..#/### => #..#/..../..../#..#"}, 2)
}

func SolveFirst() {
	solver(util.ReadTxt("day21/input.txt"), 5)
}

func SolveSecond() {
	solver(util.ReadTxt("day21/input.txt"), 18)
}
