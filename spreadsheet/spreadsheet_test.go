package spreadsheet_test

import (
	"testing"

	"github.com/baileys-am/aoc-2017/spreadsheet"
)

func TestTableIsCreatedWithExpectedRowsAndColumns(t *testing.T) {
	s := "88 90	25\r\n20	20	21"
	dels := []rune{' ', '	'}
	ss := spreadsheet.New(s, dels)

	if ss.Rows != 2 || ss.Columns != 3 {
		t.Error("failed to create table")
	}
}

func TestReturnsMinRowValue(t *testing.T) {
	s := "88 90	25\r\n20	20	21"
	dels := []rune{' ', '	'}
	ss := spreadsheet.New(s, dels)

	if ss.GetRowMin(0) != 25 || ss.GetRowMin(1) != 20 {
		t.Error("failed to get min values for each row", ss.GetRowMin(0))
	}
}

func TestReturnsMaxRowValue(t *testing.T) {
	s := "88 90	25\r\n20	20	21"
	dels := []rune{' ', '	'}
	ss := spreadsheet.New(s, dels)

	if ss.GetRowMax(0) != 90 || ss.GetRowMax(1) != 21 {
		t.Error("failed to get max values for each row")
	}
}

func TestReturnsRangeOfRowValue(t *testing.T) {
	s := "88 90	25\r\n20	20	21"
	dels := []rune{' ', '	'}
	ss := spreadsheet.New(s, dels)

	if ss.GetRowRange(0) != 65 || ss.GetRowRange(1) != 1 {
		t.Error("failed to get max values for each row")
	}
}

func TestReturnsRowDivisibleFactor(t *testing.T) {
	s := "5 9 2 8\r\n9 4 7 3\r\n3 8 6 5"
	dels := []rune{' ', '	'}
	ss := spreadsheet.New(s, dels)

	if ss.GetRowDivisibleFactor(0) != 4 || ss.GetRowDivisibleFactor(1) != 3 || ss.GetRowDivisibleFactor(2) != 2 {
		t.Error("failed to get max values for each row")
	}
}
