package day13

import (
	"strconv"
	"strings"

	. "github.com/AlbertBrand/aoc2017/util"
)

type Layers map[int]int

func solver(lines []string) {
	severity := checkSeverity(makeLayers(lines))
	println(severity)
}

func solver2(lines []string) {
	delay := 1
	for {
		caught := checkCaught(makeLayers(lines), delay)
		if !caught {
			break
		}
		delay++
	}
	println(delay)
}

func checkSeverity(layers Layers) int {
	severity := 0
	for depth, scanRange := range layers {
		// check caught
		if depth%(2*(scanRange-1)) == 0 {
			severity += depth * scanRange
		}
	}
	return severity
}

func checkCaught(layers Layers, delay int) bool {
	for depth, scanRange := range layers {
		// check caught
		if (depth+delay)%(2*(scanRange-1)) == 0 {
			return true
		}
	}
	return false
}

func makeLayers(lines []string) Layers {
	layers := make(Layers)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		pos, _ := strconv.Atoi(parts[0])
		scanRange, _ := strconv.Atoi(parts[1])
		layers[pos] = scanRange
	}
	return layers
}

func TestFirst() {
	solver(ReadTxt("day13/test.txt"))
}

func SolveFirst() {
	solver(ReadTxt("day13/input.txt"))
}

func TestSecond() {
	solver2(ReadTxt("day13/test.txt"))
}

func SolveSecond() {
	solver2(ReadTxt("day13/input.txt"))
}
