package day12

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Programs map[int][]int
type Seen map[int]bool

func readTxt(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

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
	solver(readTxt("day12/test.txt"))
}

func SolveFirst() {
	solver(readTxt("day12/input.txt"))
}

func TestSecond() {
	solver2(readTxt("day12/test.txt"))
}
func SolveSecond() {
	solver2(readTxt("day12/input.txt"))
}
