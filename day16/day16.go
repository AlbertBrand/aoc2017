package day16

import (
	"strconv"
	"strings"

	"github.com/AlbertBrand/aoc2017/util"
)

type CommandType int

const (
	Spin CommandType = iota
	Exchange
	Partner
)

type Command struct {
	ct CommandType

	start int

	a int
	b int

	x string
	y string
}

func parseLines(lines []string, max int) []Command {
	commands := make([]Command, 0)
	for _, line := range lines {
		switch line[0] {
		case 's':
			amount, _ := strconv.Atoi(line[1:])
			start := max - amount
			commands = append(commands, Command{
				ct:    Spin,
				start: start,
			})
		case 'x':
			parts := strings.Split(line[1:], "/")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			commands = append(commands, Command{
				ct: Exchange,
				a:  a,
				b:  b,
			})
		case 'p':
			parts := strings.Split(line[1:], "/")
			commands = append(commands, Command{
				ct: Partner,
				x:  parts[0],
				y:  parts[1],
			})
		}
	}
	return commands
}

func performCommands(commands []Command, programs string) string {
	for _, command := range commands {
		switch command.ct {
		case Spin:
			programs = programs[command.start:] + programs[:command.start]
		case Exchange:
			bp := []byte(programs)
			bp[command.a], bp[command.b] = bp[command.b], bp[command.a]
			programs = string(bp)
		case Partner:
			replacer := strings.NewReplacer(command.x, command.y, command.y, command.x)
			programs = replacer.Replace(programs)
		}
	}
	return programs
}

func solver(max int, lines []string, iter int) {
	programs := ""
	intA := int('a')
	for i := intA; i < intA+max; i++ {
		programs += string(i)
	}

	commands := parseLines(lines, max)

	cache := make(map[int]string)
	i := 0
	initial := programs
	for {
		cache[i] = programs
		if i != 0 && programs == initial {
			break
		} else {
			programs = performCommands(commands, programs)
		}
		i++
	}

	println(cache[iter%i])
}

func TestFirst() {
	solver(5, util.ReadAndSplit("day16/test.txt", ","), 1)
}

func SolveFirst() {
	solver(16, util.ReadAndSplit("day16/input.txt", ","), 1)
}

func SolveSecond() {
	solver(16, util.ReadAndSplit("day16/input.txt", ","), 1000000000)
}
