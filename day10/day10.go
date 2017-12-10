package day10

import "fmt"

func reverse(numbers []int, start int, size int) []int {
	numSize := len(numbers)
	for i := size/2 - 1; i >= 0; i-- {
		a := (i + start) % numSize
		b := (size - 1 - i + start) % numSize
		numbers[a], numbers[b] = numbers[b], numbers[a]
	}
	return numbers
}

func solver(count int, lengths []int) {
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = i
	}

	pos := 0
	skip := 0
	for _, length := range lengths {
		numbers = reverse(numbers, pos, length)
		//fmt.Printf("%v %d %d\n", numbers, pos, skip)
		pos += length + skip
		pos = pos % count
		skip++
	}
	println(numbers[0] * numbers[1])
}

func TestFirst() {
	solver(5, []int{3, 4, 1, 5})
}

func SolveFirst() {
	solver(256, []int{225, 171, 131, 2, 35, 5, 0, 13, 1, 246, 54, 97, 255, 98, 254, 110})
}

func solver2(input string) {
	lengths := make([]int, len(input))
	for k, v := range input {
		lengths[k] = int(v)
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	numbers := make([]int, 256)
	for i := 0; i < 256; i++ {
		numbers[i] = i
	}

	pos := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, length := range lengths {
			numbers = reverse(numbers, pos, length)
			pos += length + skip
			pos = pos % 256
			skip++
		}
	}

	hex := ""
	for i := 0; i < 16; i++ {
		result := numbers[i*16]
		for j := 1; j < 16; j++ {
			result = result ^ numbers[i*16+j]
		}

		hex += fmt.Sprintf("%02x", result)
	}
	println(hex)
}

func TestSecond() {
	solver2("")
	solver2("AoC 2017")
	solver2("1,2,3")
	solver2("1,2,4")
}

func SolveSecond() {
	solver2("225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110")
}
