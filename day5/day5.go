package day5

import (
	"bufio"
	"os"
	"strconv"
)

func readInts(filename string) []int {
	file, _ := os.Open(filename)
	defer file.Close()

	ints := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, i)
	}

	return ints
}

func solver(jumplist []int) {
	pos := 0
	nofJumps := 0
	length := len(jumplist) - 1

	for {
		jump := jumplist[pos]
		jumplist[pos]++
		pos += jump
		nofJumps++
		if pos > length {
			break
		}
	}
	println(nofJumps)
}

func solver2(jumplist []int) {
	pos := 0
	nofJumps := 0
	length := len(jumplist) - 1

	for {
		jump := jumplist[pos]
		if jump >= 3 {
			jumplist[pos]--
		} else {
			jumplist[pos]++
		}
		pos += jump
		nofJumps++
		if pos > length {
			break
		}
	}
	println(nofJumps)
}

func TestFirst() {
	solver([]int{0, 3, 0, 1, -3})
}

func SolveFirst() {
	solver(readInts("day5/input.txt"))
}

func TestSecond() {
	solver2([]int{0, 3, 0, 1, -3})
}

func SolveSecond() {
	solver2(readInts("day5/input.txt"))
}
