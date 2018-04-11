package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Get input
	instr := getInstructions()

	// Part 1 answer
	part1jumps := getPart1Ans(instr)
	fmt.Printf("Part 1, number of jumps: %d\r\n", part1jumps)

	// Part 2 answer
	part2jumps := getPart2Ans(instr)
	fmt.Printf("Part 2, number of strange jumps: %d\r\n", part2jumps)
}

// jump executes an instruction set at given position and returns new instructions and next position to execute at
func jump(instructions []int, position int) ([]int, int) {
	// Create copy of instruction set
	var newInstructions = make([]int, len(instructions))
	copy(newInstructions[:], instructions)

	offset := newInstructions[position]
	newPosition := position + offset
	newInstructions[position]++

	return newInstructions, newPosition
}

// strangJump executes an instruction set at given position and returns new instructions and next position to execute at
func strangeJump(instructions []int, position int) ([]int, int) {
	// Create copy of instruction set
	var newInstructions = make([]int, len(instructions))
	copy(newInstructions[:], instructions)

	offset := newInstructions[position]
	newPosition := position + offset
	if newInstructions[position] >= 3 {
		newInstructions[position]--
	} else {
		newInstructions[position]++
	}

	return newInstructions, newPosition
}

func getInstructions() []int {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	input := string(dat)
	strInput := strings.Split(input, "\r\n")

	instr := make([]int, 0, len(strInput))
	for _, a := range strInput {
		i, err := strconv.Atoi(a)
		check(err)
		instr = append(instr, i)
	}

	return instr
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPart1Ans(instr []int) int {
	pos := 0
	jumps := 0
	for {
		instr, pos = jump(instr, pos)
		jumps++
		if pos < 0 || pos >= len(instr) {
			break
		}
	}

	return jumps
}

func getPart2Ans(instr []int) int {
	pos := 0
	jumps := 0
	for {
		instr, pos = strangeJump(instr, pos)
		jumps++
		if pos < 0 || pos >= len(instr) {
			break
		}
	}

	return jumps
}
