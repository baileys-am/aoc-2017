package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Sort string inspired from:
// https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte

type chararray []rune

func (c chararray) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c chararray) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c chararray) Len() int {
	return len(c)
}

func SortString(s string) string {
	r := chararray(s)
	sort.Sort(r)
	return string(r)
}

func Part1ValidatePassphrase(passphrase string) bool {
	ss := strings.Split(passphrase, " ")
	words := make(map[string]bool)

	for i := range ss {
		_, exists := words[ss[i]]
		if exists {
			return false
		}

		words[ss[i]] = true
	}

	return true
}

func Part2ValidatePassphrase(passphrase string) bool {
	ss := strings.Split(passphrase, " ")
	words := make(map[string]bool)

	for i := range ss {
		sortedpp := SortString(ss[i])
		_, exists := words[sortedpp]
		if exists {
			return false
		}

		words[sortedpp] = true
	}

	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Get passphrases
	dat, err := ioutil.ReadFile("day04-passphrases.txt")
	check(err)
	input := string(dat)
	pp := strings.Split(input, "\r\n")

	// Part 1
	passing := 0
	for i := range pp {
		passes := Part1ValidatePassphrase(pp[i])
		if passes {
			passing++
		}
	}
	fmt.Printf("Part 1, number of passing passphrases: %d\r\n", passing)

	// Part 1
	passing = 0
	for i := range pp {
		passes := Part2ValidatePassphrase(pp[i])
		if passes {
			passing++
		}
	}
	fmt.Printf("Part 1, number of passing passphrases: %d\r\n", passing)
}
