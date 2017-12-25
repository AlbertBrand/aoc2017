package day24

import (
	"strings"
	"strconv"
	
	"github.com/AlbertBrand/aoc2017/util"
)

type Component struct {
	portA int
	portB int
}

type Bridge []Component

func makeCmps(lines []string) []Component {
	cmps := make([]Component, 0)
	for _, line := range lines {
		parts := strings.Split(line, "/")
		portA, _ := strconv.Atoi(parts[0])
		portB, _ := strconv.Atoi(parts[1])
		cmps = append(cmps, Component{portA, portB})
	}
	return cmps
}

func remove(cmps []Component, rem Component) []Component {
	over := make([]Component, 0)
	for _, cmp := range cmps {
		if !(rem.portA == cmp.portA &&
			rem.portB == cmp.portB) {
			over = append(over, cmp)
		}
	}
	return over
}

func find(cmps []Component, port int) []Component {
	match := make([]Component, 0)
	for _, cmp := range cmps {
		if cmp.portA == port ||
			cmp.portB == port {
			match = append(match, cmp)
		}
	}
	return match
} 

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func build(cmps []Component, port int) int {
	match := find(cmps, port)
	
	maxStrength := 0
	for _, cmp := range match {
		avail := remove(cmps, cmp)
		var newPort int
		if cmp.portA == port {
			newPort = cmp.portB
		} else {
			newPort = cmp.portA
		}
		strength := build(avail, newPort)
		maxStrength = max(maxStrength, strength)
	}
	if len(match) > 0 {
		return 2*port + maxStrength
	}
	return port
}

func build2(cmps []Component, port int) []Bridge {
	match := find(cmps, port)
	
	bridges := make([]Bridge, 0)
	for _, cmp := range match {
		avail := remove(cmps, cmp)
		var newPort int
		if cmp.portA == port {
			newPort = cmp.portB
		} else {
			newPort = cmp.portA
		}
		newBridges := build2(avail, newPort)
		for _, bridge := range newBridges {
			fullBridge := append(Bridge{}, cmp)
			fullBridge = append(fullBridge, bridge...)
			bridges = append(bridges, fullBridge)
		}
		bridges = append(bridges, Bridge{cmp})
	}
	return bridges
}

func solver(lines []string) {
	cmps := makeCmps(lines)
	
	max := build(cmps, 0)
	println(max)
}

func str(bridge Bridge) int {
	str := 0
	for _, cmp := range bridge {
		str += cmp.portA
		str += cmp.portB
	}
	return str
}

func solver2(lines []string) {
	cmps := makeCmps(lines)
	
	bridges := build2(cmps, 0)
	maxLen := 0
	maxStr := 0
	for _, bridge := range bridges {
		if len(bridge) > maxLen {
			maxLen = max(len(bridge), maxLen)
			maxStr = str(bridge)
		}
		if len(bridge) == maxLen {
			maxStr = max(str(bridge), maxStr)
		}
	}
	println(maxLen, maxStr)
}

func TestFirst() {
	solver(util.ReadTxt("day24/test.txt"))
}

func SolveFirst() {
	solver(util.ReadTxt("day24/input.txt"))
}

func TestSecond() {
	solver2(util.ReadTxt("day24/test.txt"))
}

func SolveSecond() {
	solver2(util.ReadTxt("day24/input.txt"))
}
