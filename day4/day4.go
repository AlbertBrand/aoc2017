package day4

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readTxt(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return []string{}
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return []string{}
	}

	return lines
}

func isValid(passphrase string) bool {
	parts := strings.Split(passphrase, " ")
	partMap := make(map[string]bool)
	for _, part := range parts {
		_, exists := partMap[part]
		if exists {
			return false
		}
		partMap[part] = true
	}
	return true
}

func isValid2(passphrase string) bool {
	parts := strings.Split(passphrase, " ")
	partMap := make(map[string]bool)
	for _, part := range parts {
		// sort chars in string
		chars := make([]int, 0)
		for _, char := range part {
			chars = append(chars, int(char))
		}
		sort.Ints(chars)
		part := ""
		for _, char := range chars {
			part = part + string(char)
		}

		_, exists := partMap[part]
		if exists {
			return false
		}
		partMap[part] = true
	}
	return true
}

func solver(passphrases []string, isValid func(pp string) bool) int {
	valid := 0
	for _, passphrase := range passphrases {
		if isValid(passphrase) {
			valid++
		}
	}
	return valid
}

func TestFirst() {
	println(isValid("aa bb cc dd ee"))
	println(isValid("aa bb cc dd aa"))
	println(isValid("aa bb cc dd aaa"))
}

func SolveFirst() {
	println(solver(readTxt("day4/input.txt"), isValid))
}

func TestSecond() {
	println(isValid2("abcde fghij"))
	println(isValid2("abcde xyz ecdab"))
	println(isValid2("a ab abc abd abf abj"))
	println(isValid2("iiii oiii ooii oooi oooo"))
	println(isValid2("oiii ioii iioi iiio"))
}

func SolveSecond() {
	println(solver(readTxt("day4/input.txt"), isValid2))
}
