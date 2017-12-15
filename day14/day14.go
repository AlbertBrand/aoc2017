package day14

import (
	"fmt"

	"github.com/AlbertBrand/aoc2017/day10"
)

func countBits(num int) int {
	if num > 0 {
		return (num % 2) + countBits(num>>1)
	}
	return 0
}

func solver(input string) {
	total := 0
	for i := 0; i < 128; i++ {
		row := fmt.Sprintf("%s-%d", input, i)
		hash := day10.KnotHash(row)
		for _, i := range hash {
			total += countBits(i)
		}
	}
	println(total)
}

type Grid [128][128]bool

func (grid *Grid) clear(i int, j int) {
	if !grid[i][j] {
		return
	}
	grid[i][j] = false
	if i > 0 {
		grid.clear(i-1, j)
	}
	if j > 0 {
		grid.clear(i, j-1)
	}
	if i < 127 {
		grid.clear(i+1, j)
	}
	if j < 127 {
		grid.clear(i, j+1)
	}
}

func solver2(input string) {
	var grid Grid
	for i := 0; i < 128; i++ {
		row := fmt.Sprintf("%s-%d", input, i)
		hash := day10.KnotHash(row)
		for k, j := range hash {
			bitpos := 7
			for j > 0 {
				grid[i][k*8+bitpos] = j%2 == 1
				j = j >> 1
				bitpos--
			}
		}
	}
	regions := 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if grid[i][j] {
				grid.clear(i, j)
				regions++
			}
		}
	}
	println(regions)
}

func TestFirst() {
	solver("flqrgnkx")
}

func SolveFirst() {
	solver("uugsqrei")
}

func TestSecond() {
	solver2("flqrgnkx")
}

func SolveSecond() {
	solver2("uugsqrei")
}
