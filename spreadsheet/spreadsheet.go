package spreadsheet

import (
	"math"
	"strconv"
	"strings"
)

// Spreadsheet provides utility to analyze a numeric table
type Spreadsheet struct {
	Rows    int
	Columns int
	table   [][]string
}

// New creates a new spreadsheet from csv with custom delimeters
func New(spreadsheet string, delimeters []rune) Spreadsheet {
	rowStrings := strings.Split(spreadsheet, "\r\n")
	table := make([][]string, len(rowStrings))
	maxColumns := 0
	for i := range table {
		table[i] = strings.FieldsFunc(rowStrings[i], func(r rune) bool {
			for del := range delimeters {
				if r == delimeters[del] {
					return true
				}
			}
			return false
		})

		if len(table[i]) > maxColumns {
			maxColumns = len(table[i])
		}
	}

	s := Spreadsheet{len(table), maxColumns, table}
	return s
}

// GetRowMin returns the minimum value in a spreadsheet row.
// Returns 0 if no numeric values are found.
func (s Spreadsheet) GetRowMin(row int) float64 {
	min := float64(0)
	first := true
	for c := range s.table[row] {
		num, err := strconv.ParseFloat(s.table[row][c], 10)
		if err == nil && (first || num < min) {
			min = num
			first = false
		}
	}

	return min
}

// GetRowMax returns the maximum value in a spreadsheet row.
// Returns 0 if no numeric values are found.
func (s Spreadsheet) GetRowMax(row int) float64 {
	max := float64(0)
	first := true
	for c := range s.table[row] {
		num, err := strconv.ParseFloat(s.table[row][c], 10)
		if err == nil && (first || num > max) {
			max = num
			first = false
		}
	}

	return max
}

// GetRowRange returns the maximum value in a spreadsheet row.
// Returns 0 if no numeric values are found.
func (s Spreadsheet) GetRowRange(row int) float64 {
	max := s.GetRowMax(row)
	min := s.GetRowMin(row)

	return math.Abs(max - min)
}

// GetRowDivisibleFactor returns the first divisble factor between row elements
// Returns 0 if no numeric values or factor is found.
func (s Spreadsheet) GetRowDivisibleFactor(row int) int {
	for c := range s.table[row] {
		num1, err := strconv.ParseInt(s.table[row][c], 10, 32)
		if err == nil {
			for i := range s.table[row] {
				if c != i {
					num2, err := strconv.ParseInt(s.table[row][i], 10, 32)
					if err == nil && num1%num2 == 0 {
						return int(num1 / num2)
					}
				}
			}
		}
	}

	return 0
}
