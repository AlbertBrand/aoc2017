package day18

import (
	"strconv"
	"strings"

	"github.com/AlbertBrand/aoc2017/util"
)

func solver(lines []string) {
	instructions := make([]func(), 0)
	registers := make(map[string]int)
	pointer := 0
	sound := 0
	quit := false

	getValue := func(value int, valreg string) int {
		if valreg != "" {
			return registers[valreg]
		}
		return value
	}
	snd := func(register string) func() {
		return func() {
			sound = registers[register]
			pointer++
		}
	}
	set := func(register string, value int, valreg string) func() {
		return func() {
			registers[register] = getValue(value, valreg)
			pointer++
		}
	}
	add := func(register string, value int, valreg string) func() {
		return func() {
			registers[register] += getValue(value, valreg)
			pointer++
		}
	}
	mul := func(register string, value int, valreg string) func() {
		return func() {
			registers[register] *= getValue(value, valreg)
			pointer++
		}
	}
	mod := func(register string, value int, valreg string) func() {
		return func() {
			registers[register] %= getValue(value, valreg)
			pointer++
		}
	}
	rcv := func(register string) func() {
		return func() {
			if registers[register] != 0 {
				println(sound)
				quit = true
			}
			pointer++
		}
	}
	jgz := func(register string, value int, valreg string) func() {
		return func() {
			if valreg != "" {
				value = registers[valreg]
			}
			if registers[register] > 0 {
				pointer += value
			} else {
				pointer++
			}
		}
	}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		instruction := parts[0]
		// register is bugged but for solution 1 it works
		register := parts[1]
		value := 0
		valreg := ""
		if len(parts) > 2 {
			var err error
			value, err = strconv.Atoi(parts[2])
			if err != nil {
				valreg = parts[2]
			}
		}

		switch instruction {
		case "snd":
			instructions = append(instructions, snd(register))
		case "set":
			instructions = append(instructions, set(register, value, valreg))
		case "add":
			instructions = append(instructions, add(register, value, valreg))
		case "mul":
			instructions = append(instructions, mul(register, value, valreg))
		case "mod":
			instructions = append(instructions, mod(register, value, valreg))
		case "rcv":
			instructions = append(instructions, rcv(register))
		case "jgz":
			instructions = append(instructions, jgz(register, value, valreg))
		}
	}

	maxpos := len(instructions) - 1
	for {
		instructions[pointer]()
		if quit || pointer < 0 || pointer > maxpos {
			break
		}
	}
}

type Command int

const (
	snd Command = iota
	rcv
	set
	add
	mul
	mod
	jgz
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
	id           int
	instructions []Instruction
	registers    map[string]int
	pointer      int
	sndChan      chan<- int
	rcvChan      <-chan int
	sendCnt      int
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
	end := false

	switch i.command {
	case snd:
		m.sndChan <- m.getValue(i.arg0)
		m.sendCnt++
	case rcv:
		select {
		case value := <-m.rcvChan:
			m.registers[i.arg0.register] = value
		default:
			nextPointer = m.pointer
			end = true
		}
	case set:
		m.registers[i.arg0.register] = m.getValue(i.arg1)
	case add:
		m.registers[i.arg0.register] += m.getValue(i.arg1)
	case mul:
		m.registers[i.arg0.register] *= m.getValue(i.arg1)
	case mod:
		m.registers[i.arg0.register] %= m.getValue(i.arg1)
	case jgz:
		if m.getValue(i.arg0) > 0 {
			nextPointer = m.pointer + m.getValue(i.arg1)
		}
	}

	m.pointer = nextPointer
	return end
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
	case "snd":
		return Instruction{command: snd, arg0: arg0}
	case "rcv":
		return Instruction{command: rcv, arg0: arg0}
	case "set":
		return Instruction{set, arg0, arg1}
	case "add":
		return Instruction{add, arg0, arg1}
	case "mul":
		return Instruction{mul, arg0, arg1}
	case "mod":
		return Instruction{mod, arg0, arg1}
	case "jgz":
		return Instruction{jgz, arg0, arg1}
	default:
		panic("unknown instruction")
	}
}

func createMachine(lines []string, id int, sndChan chan<- int, rcvChan <-chan int) *Machine {
	m := Machine{
		id:           id,
		instructions: make([]Instruction, 0),
		registers:    make(map[string]int),
		pointer:      0,
		sndChan:      sndChan,
		rcvChan:      rcvChan,
		sendCnt:      0,
	}
	m.registers["p"] = id

	for _, line := range lines {
		m.instructions = append(m.instructions, parseInstruction(line))
	}

	return &m
}

func solver2(lines []string) {
	chan0 := make(chan int, 1024)
	chan1 := make(chan int, 1024)
	p0 := createMachine(lines, 0, chan0, chan1)
	p1 := createMachine(lines, 1, chan1, chan0)

	exec1 := false
	for {
		if exec1 {
			if end := p1.step(); end {
				if len(chan1) == 0 {
					break
				}
				exec1 = false
			}
		} else {
			if end := p0.step(); end {
				if len(chan0) == 0 {
					break
				}
				exec1 = true
			}
		}
	}
	println(p1.sendCnt)
}

func TestFirst() {
	solver(util.ReadTxt("day18/test.txt"))
}

func SolveFirst() {
	solver(util.ReadTxt("day18/input.txt"))
}

func TestSecond() {
	solver2(util.ReadTxt("day18/test2.txt"))
}

func SolveSecond() {
	solver2(util.ReadTxt("day18/input.txt"))
}
