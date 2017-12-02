package day2

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readTxt(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return [][]string{}
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return [][]string{}
	}

	return records
}

func solver(records [][]string) {
	checksum := 0
	for _, record := range records {
		vals := make([]int, 0)
		for _, val := range record {
			ival, _ := strconv.Atoi(val)
			vals = append(vals, ival)
		}

		sort.Ints(vals)
		rmin := vals[0]
		rmax := vals[len(vals)-1]

		rcheck := rmax - rmin
		checksum += rcheck
	}
	println(checksum)
}

func solver2(records [][]string) {
	checksum := 0
	for _, record := range records {
		for _, val := range record {
			ival, _ := strconv.Atoi(val)
			fval := float64(ival)
			for _, val2 := range record {
				ival2, _ := strconv.Atoi(val2)
				if ival2 == ival {
					continue
				}
				fval2 := float64(ival2)
				div := fval / fval2
				if div == float64(int(div)) {
					checksum += int(div)
					break
				}
			}
		}
	}
	println(checksum)
}

func TestFirst() {
	solver(readTxt("day2/test1.txt"))
}

func SolveFirst() {
	solver(readTxt("day2/input.txt"))
}

func TestSecond() {
	solver2(readTxt("day2/test2.txt"))
}

func SolveSecond() {
	solver2(readTxt("day2/input.txt"))
}
