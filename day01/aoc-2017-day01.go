package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	// Get sequence
	dat, err := ioutil.ReadFile("day01-sequence.txt")
	check(err)
	input := string(dat)

	// Part 1
	var captcha1 = dupsum(input, 1)
	fmt.Println(captcha1)

	// Part 2
	var captcha2 = dupsum(input, len(input)/2)
	fmt.Println(captcha2)
}

func dupsum(input string, dupDist int) (sum int) {
	var length = len(input)
	for i := 0; i < length; i++ {
		if (i+dupDist < length && input[i] == input[i+dupDist]) || (i+dupDist >= length && input[i] == input[i+dupDist-length]) {
			num, err := strconv.ParseInt(string(input[i]), 10, 32)
			check(err)
			sum += int(num)
		}
	}

	return sum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
