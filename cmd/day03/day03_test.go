package main

import (
	"strconv"
	"testing"
)

/*
Testing:
[*] Grid size detection
[*] Grid perimeter calculation
[*] Calculate wrap around perimeter from bottom right cell
[*] Calculate X and Y using wrap
[*] Calculate Manhattan distance
*/

func TestGridSizeDetection(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 1},
		{2, 3},
		{10, 5},
		{26, 7},
		{52, 9},
	}

	for _, tt := range tests {
		actual := GetGridDim(tt.in)

		if tt.out != actual {
			t.Error("failed detect grid size")
		}
	}
}

func TestGridPerimeter(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 0},
		{3, 8},
		{5, 16},
		{7, 24},
		{9, 32},
	}

	for _, tt := range tests {
		actual := GetGridPerimeter(tt.in)

		if tt.out != actual {
			t.Error("failed to calculate grid perimeter")
		}
	}
}

func TestGridWrapCount(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 0},
		{2, 7},
		{10, 15},
		{26, 23},
		{52, 29},
	}

	for _, tt := range tests {
		actual := GetGridWrap(tt.in)

		if tt.out != actual {
			t.Error("failed to calculate grid wrap")
		}
	}
}

func TestGridGetXY(t *testing.T) {
	tests := []struct {
		in   int
		outX int
		outY int
	}{
		{1, 0, 0},
		{2, 1, 0},
		{10, 2, -1},
		{13, 2, 2},
		{14, 1, 2},
		{16, -1, 2},
		{17, -2, 2},
		{23, 0, -2},
		{25, 2, -2},
		{26, 3, -2},
		{52, 4, -1},
		{68, -4, 1},
		{72, -4, -3},
		{74, -3, -4},
	}

	for _, tt := range tests {
		actualX, actualY := GetGridXY(tt.in)

		if tt.outX != actualX || tt.outY != actualY {
			t.Error("failed to calculate grid XY (" + strconv.Itoa(actualX) + ", " + strconv.Itoa(actualY) + ")")
		}
	}
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for _, tt := range tests {
		actual := GetGridManhattanDistance(tt.in)

		if tt.out != actual {
			t.Error("failed to calculate Manhattan distance")
		}
	}
}

func TestSpiralSum(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 5},
		{6, 10},
		{17, 147},
	}

	for _, tt := range tests {
		actual := GetSpiralSumValue(tt.in)

		if tt.out != actual {
			t.Errorf("failed to calculate spiral sum value. Actual: %d", actual)
		}
	}
}
