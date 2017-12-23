package day23

import (
	"strconv"
	"strings"

	"github.com/AlbertBrand/aoc2017/util"
)

type Command int

const (
	set Command = iota
	sub
	mul
	jnz
)

type Arg struct {
	register string
	value    int
}

type Instruction struct {
	command Command
	arg0    Arg
	arg1    Arg
}

type Machine struct {
	instructions []Instruction
	registers    map[string]int
	pointer      int
	mulCnt       int
}

func (m *Machine) getValue(a Arg) int {
	if a.register != "" {
		return m.registers[a.register]
	}
	return a.value
}

func (m *Machine) step() bool {
	if m.pointer >= len(m.instructions) || m.pointer < 0 {
		return true
	}

	i := m.instructions[m.pointer]
	nextPointer := m.pointer + 1

	switch i.command {
	case set:
		m.registers[i.arg0.register] = m.getValue(i.arg1)
	case sub:
		m.registers[i.arg0.register] -= m.getValue(i.arg1)
	case mul:
		m.registers[i.arg0.register] *= m.getValue(i.arg1)
		m.mulCnt++
	case jnz:
		if m.getValue(i.arg0) != 0 {
			nextPointer = m.pointer + m.getValue(i.arg1)
		}
	}

	m.pointer = nextPointer
	return false
}

func parseArg(arg string) Arg {
	value, err := strconv.Atoi(arg)
	if err != nil {
		return Arg{register: arg}
	}
	return Arg{value: value}
}

func parseInstruction(line string) Instruction {
	parts := strings.Split(line, " ")

	instruction := parts[0]
	arg0 := parseArg(parts[1])
	var arg1 Arg
	if len(parts) > 2 {
		arg1 = parseArg(parts[2])
	}

	switch instruction {
	case "set":
		return Instruction{set, arg0, arg1}
	case "sub":
		return Instruction{sub, arg0, arg1}
	case "mul":
		return Instruction{mul, arg0, arg1}
	case "jnz":
		return Instruction{jnz, arg0, arg1}
	default:
		panic("unknown instruction")
	}
}

func makeMachine(lines []string) Machine {
	m := Machine{
		instructions: make([]Instruction, 0),
		registers:    make(map[string]int),
		pointer:      0,
	}
	for _, line := range lines {
		m.instructions = append(m.instructions, parseInstruction(line))
	}

	return m
}

func solver(lines []string) {
	m := makeMachine(lines)

	for {
		if end := m.step(); end {
			break
		}
	}

	println(m.mulCnt)
}

func solver2() {
	h := 0
	b := 81
	c := b
	b = b * 100
	b = b + 100000
	c = b + 17000

main:
	for {
		f := 1
		d := 2

		for {
			if b%d == 0 {
				f = 0
			}
			d++
			if d != b {
				continue
			}
			if f == 0 {
				h++
			}
			if b == c {
				break main
			}
			b += 17
			break
		}
	}

	print(h)
}

func SolveFirst() {
	solver(util.ReadTxt("day23/input.txt"))
}

func SolveSecond() {
	solver2()
}
