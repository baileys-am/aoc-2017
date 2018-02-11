package main

import (
	"fmt"
	"io/ioutil"

	"github.com/baileys-am/aoc-2017/spreadsheet"
)

func main() {
	// Get sequence
	dat, err := ioutil.ReadFile("day02-spreadsheet.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	dels := []rune{' ', '	'}
	ss := spreadsheet.New(input, dels)

	// Part 1
	checksum := 0
	for i := 0; i < ss.Rows; i++ {
		checksum += int(ss.GetRowRange(i))
	}

	fmt.Println("checksum is", checksum)

	// Part 2
	factorSum := 0
	for i := 0; i < ss.Rows; i++ {
		factorSum += int(ss.GetRowDivisibleFactor(i))
	}

	fmt.Println("factor sum is", factorSum)
}
