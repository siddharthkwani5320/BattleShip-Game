package component

import (
	"fmt"
	"os"
	"testing"
)

type TestBoard struct {
	row int
	col int
}

var testboard = []TestBoard{
	{0, 0},
	{4, 4},
	{5, 3},
	{5, 2},
	{5, 5},
	{6, 6},
	{7, 8},
	{13, 12},
	{10, 9},
}

func TestNewBoard(t *testing.T) {
	for _, val := range testboard {
		board, err := NewBoard(val.row, val.col)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		expectedRow := val.row
		expectedCol := val.col
		row := board.GetRow()
		col := board.GetCol()
		if row != expectedRow || col != expectedCol {
			fmt.Println("Result does not match with expected ")
		}
	}
}

func TestGetRow(t *testing.T) {
	for _, val := range testboard {
		board, err := NewBoard(val.row, val.col)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		expectedRow := val.row
		row := board.GetRow()
		if row != expectedRow {
			fmt.Println("Result does not match with expected ")
		}
	}
}

func TestGetCol(t *testing.T) {
	for _, val := range testboard {
		board, err := NewBoard(val.row, val.col)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		expectedCol := val.col
		col := board.GetCol()
		if col != expectedCol {
			fmt.Println("Result does not match with expected ")
		}
	}
}
