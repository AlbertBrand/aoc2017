package day6

import (
	"strconv"
	"strings"
)

func largestBlock(blocks []int) (int, int) {
	pos := 0
	max := 0
	for k, block := range blocks {
		if block > max {
			pos = k
			max = block
		}
	}
	return pos, max
}

func hasher(blocks []int) string {
	b := make([]string, len(blocks))
	for i, v := range blocks {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, "|")
}

func solver(blocks []int) {
	entries := make([]string, 0)
	cycle := 0

	for {
		length := len(blocks)
		pos, max := largestBlock(blocks)
		spread := max / length
		remainder := max % length

		blocks[pos] = 0
		for i := 0; i < len(blocks); i++ {
			pos++
			pos = pos % length
			blocks[pos] += spread
			if i < remainder {
				blocks[pos] += 1
			}
		}
		cycle++

		hash := hasher(blocks)
		for k, entry := range entries {
			if entry == hash {
				println("cycles", cycle)
				println("loop size", len(entries)-k)
				return
			}
		}
		entries = append(entries, hash)
	}
}

func Test() {
	solver([]int{0, 2, 7, 0})
}

func Solve() {
	solver([]int{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6})
}
