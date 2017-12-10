package day9

import (
	"io/ioutil"
)

func readFile(filename string) string {
	b, _ := ioutil.ReadFile(filename)
	return string(b)
}

func solver(stream string) {
	score := 0
	garbageCount := 0
	groupsOpen := 0
	garbageOpen := false
	ignore := false

	pos := 0
	max := len(stream) - 1
	for {
		char := stream[pos]
		if !ignore {
			if !garbageOpen {
				switch char {
				case '{':
					groupsOpen++
				case '}':
					score += groupsOpen
					groupsOpen--
				case '<':
					garbageOpen = true
				case '!':
					ignore = true
				}
			} else {
				switch char {
				case '>':
					garbageOpen = false
				case '!':
					ignore = true
				default:
					garbageCount++
				}
			}
		} else {
			ignore = false
		}

		pos++
		if pos > max {
			break
		}
	}
	println(score, garbageCount)
}

func Test() {
	//solver("{}")
	//solver("{{{}}}")
	//solver("{{},{}}")
	//solver("{{{},{},{{}}}}")
	//solver("{<a>,<a>,<a>,<a>}")
	//solver("{{<ab>},{<ab>},{<ab>},{<ab>}}")
	//solver("{{<!!>},{<!!>},{<!!>},{<!!>}}")
	//solver("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	//solver("<>")
	//solver("<random characters>")
	//solver("<<<<>")
	//solver("<{!>}>")
	//solver("<!!>")
	//solver("<!!!>>")
	//solver(`<{o"i!a,<{i<a>`)
	solver("<>")
	solver("<random characters>")
	solver("<<<<>")
	solver("<{!>}>")
	solver("<!!>")
	solver(`<{o"i!a,<{i<a>`)

}

func Solve() {
	solver(readFile("day9/input.txt"))
}
