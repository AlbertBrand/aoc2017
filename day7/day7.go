package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "github.com/AlbertBrand/aoc2017/util"
)

type Node struct {
	name       string
	weight     int
	sumWeight  int
	childNames []string
	parent     *Node
}

var re = regexp.MustCompile(`(\w+) \((\d+)\)( -> (.*))?`)

func parseLine(line string) *Node {
	result := re.FindStringSubmatch(line)
	weight, _ := strconv.Atoi(result[2])
	n := Node{
		name:   result[1],
		weight: weight,
	}
	if result[4] != "" {
		n.childNames = strings.Split(result[4], ", ")
	}
	return &n
}

func solver(lines []string) {
	nodeMap := make(map[string]*Node, 0)
	// create map
	for _, line := range lines {
		node := parseLine(line)
		nodeMap[node.name] = node
	}
	// link to parent, building a tree
	for _, node := range nodeMap {
		for _, childName := range node.childNames {
			nodeMap[childName].parent = node
		}
	}
	// start somewhere in the tree
	var curr *Node
	for _, node := range nodeMap {
		curr = node
		break
	}
	// traverse up to parent
	for {
		if curr.parent == nil {
			println(curr.name)
			break
		}
		curr = curr.parent
	}
}

func checkWeights(nodeMap map[string]*Node, node *Node) int {
	sumWeight := node.weight
	for _, childName := range node.childNames {
		child := nodeMap[childName]
		childWeight := checkWeights(nodeMap, child)
		sumWeight += childWeight
	}

	if len(node.childNames) > 1 {
		firstChild := nodeMap[node.childNames[0]]
		for _, childName := range node.childNames {
			child := nodeMap[childName]
			if child.sumWeight != firstChild.sumWeight {
				println(child.weight - (child.sumWeight - firstChild.sumWeight))
			}
		}
	}

	node.sumWeight = sumWeight
	return sumWeight
}

func solver2(lines []string) {
	nodeMap := make(map[string]*Node, 0)
	// create map
	for _, line := range lines {
		node := parseLine(line)
		nodeMap[node.name] = node
	}
	// link to parent, building a tree
	for _, node := range nodeMap {
		for _, childName := range node.childNames {
			nodeMap[childName].parent = node
		}
	}
	// start from root, check weights
	checkWeights(nodeMap, nodeMap["qibuqqg"])
}

func TestFirst() {
	fmt.Printf("%v", parseLine("gbyvdfh (155) -> xqmnq, iyoqt, dimle"))
	fmt.Printf("%v", parseLine("oweiea (97)"))
}

func SolveFirst() {
	solver(ReadTxt("day7/input.txt"))
}

func SolveSecond() {
	solver2(ReadTxt("day7/input.txt"))
}
