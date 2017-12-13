package day12

import (
	"regexp"
	"strconv"
	"strings"

	. "github.com/AlbertBrand/aoc2017/util"
)

type Programs map[int][]int
type Seen map[int]bool

var re = regexp.MustCompile(`(\d+) <-> (.*)`)

func makeProgramMap(lines []string) Programs {
	programs := make(Programs)
	for _, line := range lines {
		result := re.FindStringSubmatch(line)
		id, _ := strconv.Atoi(result[1])
		connections := make([]int, 0)
		for _, val := range strings.Split(result[2], ", ") {
			id, _ := strconv.Atoi(val)
			connections = append(connections, id)
		}
		programs[id] = connections
	}
	return programs
}

func traverseSeen(connections []int, programs Programs, seen Seen) {
	for _, val := range connections {
		_, exists := seen[val]
		if !exists {
			seen[val] = true
			traverseSeen(programs[val], programs, seen)
		}
	}
}

func solver(lines []string) {
	programs := makeProgramMap(lines)
	seen := make(Seen)
	traverseSeen([]int{0}, programs, seen)
	println(len(seen))
}

func solver2(lines []string) {
	programs := makeProgramMap(lines)
	seen := make(Seen)
	count := 0
	for id := range programs {
		_, exists := seen[id]
		if !exists {
			traverseSeen([]int{id}, programs, seen)
			count++
		}
	}
	println(count)
}

func TestFirst() {
	solver(ReadTxt("day12/test.txt"))
}

func SolveFirst() {
	solver(ReadTxt("day12/input.txt"))
}

func TestSecond() {
	solver2(ReadTxt("day12/test.txt"))
}
func SolveSecond() {
	solver2(ReadTxt("day12/input.txt"))
}
