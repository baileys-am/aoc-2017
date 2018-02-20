package main

import (
	"fmt"
	"math"
)

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

func main() {
	dist := GetGridManhattanDistance(312051)
	fmt.Printf("Part 1: Distance is %d", dist)
}
