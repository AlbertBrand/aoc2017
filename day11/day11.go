package day11

import (
	"io/ioutil"
	"strings"
)

type Vector struct {
	x, y, z int
}

var directions = map[string]Vector{
	"n":  Vector{0, 1, -1},
	"ne": Vector{1, 0, -1},
	"se": Vector{1, -1, 0},
	"s":  Vector{0, -1, 1},
	"sw": Vector{-1, 0, 1},
	"nw": Vector{-1, 1, 0},
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (c *Vector) dist() int {
	return (abs(c.x) + abs(c.y) + abs(c.z)) / 2
}

func (c *Vector) add(a Vector) {
	c.x += a.x
	c.y += a.y
	c.z += a.z
}

func solver(steps []string) {
	pos := Vector{0, 0, 0}
	var max Vector
	for _, step := range steps {
		vect, ok := directions[step]
		if !ok {
			panic("incorrect input")
		}
		pos.add(vect)
		if pos.dist() > max.dist() {
			max = pos
		}
	}
	println(pos.dist(), max.dist())
}

func Test() {
	solver([]string{"ne", "ne", "ne"})
	solver([]string{"ne", "ne", "sw", "sw"})
	solver([]string{"ne", "ne", "s", "s"})
	solver([]string{"se", "sw", "se", "sw", "sw"})
}

func readFile(filename string) []string {
	b, _ := ioutil.ReadFile(filename)
	return strings.Split(string(b), ",")
}

func Solve() {
	solver(readFile("day11/input.txt"))
}
