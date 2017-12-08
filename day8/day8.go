package day8

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Operator int

const (
	increase Operator = iota
	decrease

	larger
	largerEqual
	equal
	lesserEqual
	lesser
	notEqual
)

type Instruction struct {
	reg   string
	regOp Operator
	val   int
	ifReg string
	ifOp  Operator
	ifVal int
}

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

func parseLine(line string) Instruction {
	parts := strings.Split(line, " ")

	var regOp Operator
	switch parts[1] {
	case "inc":
		regOp = increase
	case "dec":
		regOp = decrease
	}

	val, _ := strconv.Atoi(parts[2])

	var ifOp Operator
	switch parts[5] {
	case ">":
		ifOp = larger
	case ">=":
		ifOp = largerEqual
	case "==":
		ifOp = equal
	case "<=":
		ifOp = lesserEqual
	case "<":
		ifOp = lesser
	case "!=":
		ifOp = notEqual
	}

	ifVal, _ := strconv.Atoi(parts[6])

	return Instruction{
		reg:   parts[0],
		regOp: regOp,
		val:   val,
		ifReg: parts[4],
		ifOp:  ifOp,
		ifVal: ifVal,
	}
}

var regMap = make(map[string]int)
var maxRegVal int

func apply(inst Instruction) {
	switch inst.regOp {
	case increase:
		regMap[inst.reg] += inst.val
	case decrease:
		regMap[inst.reg] -= inst.val
	}
	if regMap[inst.reg] > maxRegVal {
		maxRegVal = regMap[inst.reg]
	}
}

func solver(lines []string) {
	for _, line := range lines {
		inst := parseLine(line)
		readReg := regMap[inst.ifReg]
		switch inst.ifOp {
		case larger:
			if readReg > inst.ifVal {
				apply(inst)
			}
		case largerEqual:
			if readReg >= inst.ifVal {
				apply(inst)
			}
		case equal:
			if readReg == inst.ifVal {
				apply(inst)
			}
		case lesserEqual:
			if readReg <= inst.ifVal {
				apply(inst)
			}
		case lesser:
			if readReg < inst.ifVal {
				apply(inst)
			}
		case notEqual:
			if readReg != inst.ifVal {
				apply(inst)
			}
		}
	}
	var maxReg string
	var maxVal int
	for reg, val := range regMap {
		if val > maxVal {
			maxReg = reg
			maxVal = val
		}
	}
	println(maxReg, maxVal, maxRegVal)
}

func Test() {
	solver(readTxt("day8/test.txt"))
}

func Solve() {
	solver(readTxt("day8/input.txt"))
}
