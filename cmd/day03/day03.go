package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) northwest() Coordinate {
	return Coordinate{c.X - 1, c.Y + 1}
}

func (c *Coordinate) north() Coordinate {
	return Coordinate{c.X, c.Y + 1}
}

func (c *Coordinate) northeast() Coordinate {
	return Coordinate{c.X + 1, c.Y + 1}
}

func (c *Coordinate) east() Coordinate {
	return Coordinate{c.X + 1, c.Y}
}

func (c *Coordinate) southeast() Coordinate {
	return Coordinate{c.X + 1, c.Y - 1}
}

func (c *Coordinate) south() Coordinate {
	return Coordinate{c.X, c.Y - 1}
}

func (c *Coordinate) southwest() Coordinate {
	return Coordinate{c.X - 1, c.Y - 1}
}

func (c *Coordinate) west() Coordinate {
	return Coordinate{c.X - 1, c.Y}
}

func GetGridDim(value int) int {
	dim := int(math.Ceil(math.Sqrt(float64(value))))
	if dim%2 == 0 {
		dim++
	}

	return dim
}

func GetGridPerimeter(dim int) int {
	return 4*dim - 4
}

func GetGridWrap(value int) int {
	dim := GetGridDim(value)

	return dim*dim - value
}

func GetGridXY(value int) (x int, y int) {
	dim := GetGridDim(value)
	perimeter := GetGridPerimeter(dim)
	wrap := GetGridWrap(value)
	wrapPercentage := float64(wrap) / float64(perimeter)
	switch {
	case wrapPercentage == 0:
		x = (dim - 1) / 2
		y = -(dim - 1) / 2
		break
	case wrapPercentage < 0.25:
		x = (dim-1)/2 - wrap
		y = -(dim - 1) / 2
		break
	case wrapPercentage == 0.25:
		x = -(dim - 1) / 2
		y = -(dim - 1) / 2
		break
	case wrapPercentage > 0.25 && wrapPercentage < 0.5:
		x = -(dim - 1) / 2
		y = (dim-1)/2 - (perimeter/2 - wrap)
		break
	case wrapPercentage == 0.5:
		x = -(dim - 1) / 2
		y = (dim - 1) / 2
		break
	case wrapPercentage > 0.5 && wrapPercentage < 0.75:
		x = (wrap - perimeter/2) - (dim-1)/2
		y = (dim - 1) / 2
		break
	case wrapPercentage == 0.75:
		x = (dim - 1) / 2
		y = (dim - 1) / 2
		break
	case wrapPercentage > 0.75:
		x = (dim - 1) / 2
		y = (perimeter - wrap) - (dim-1)/2
		break
	}

	return x, y
}

func GetGridManhattanDistance(value int) int {
	x, y := GetGridXY(value)

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func GetSpiralSumValue(value int) int {
	// Initialize sum values (tracks matrix coord values)
	sumValues := make(map[Coordinate]int)
	sumValues[Coordinate{0, 0}] = 1

	// Populate values up to given value
	for i := 2; i <= value; i++ {
		// Get coordinate for value and initialize sum
		x, y := GetGridXY(i)
		coord := Coordinate{x, y}
		sumValue := 0

		// Add north west value
		nwVal, ok := sumValues[coord.northwest()]
		if ok {
			sumValue += nwVal
		}

		// Add north value
		nVal, ok := sumValues[coord.north()]
		if ok {
			sumValue += nVal
		}

		// Add north east value
		neVal, ok := sumValues[coord.northeast()]
		if ok {
			sumValue += neVal
		}

		// Add east value
		eVal, ok := sumValues[coord.east()]
		if ok {
			sumValue += eVal
		}

		// Add south east value
		seVal, ok := sumValues[coord.southeast()]
		if ok {
			sumValue += seVal
		}

		// Add south value
		sVal, ok := sumValues[coord.south()]
		if ok {
			sumValue += sVal
		}

		// Add south west value
		swVal, ok := sumValues[coord.southwest()]
		if ok {
			sumValue += swVal
		}

		// Add west value
		wVal, ok := sumValues[coord.west()]
		if ok {
			sumValue += wVal
		}

		sumValues[coord] = sumValue
	}

	// Get and return sum value with coord for given value
	x, y := GetGridXY(value)
	coord := Coordinate{x, y}
	return sumValues[coord]
}

func main() {
	input := 312051

	// Part 1
	dist := GetGridManhattanDistance(input)
	fmt.Printf("Part 1: Distance is %d\r\n", dist)

	// Part 2
	for i := 1; true; i++ {
		sumValue := GetSpiralSumValue(i)
		if sumValue > input {
			fmt.Printf("Part 2: Larger value is %d", sumValue)
			break
		}
	}
}
