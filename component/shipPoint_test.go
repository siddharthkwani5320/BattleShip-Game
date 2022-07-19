package component

import (
	"fmt"
	"testing"
)

type Points struct {
	row int
	col int
}

var testpoints = []Points{
	{4, 5},
	{2, 2},
	{4, 4},
	{0, 0},
	{6, 6},
	{0, 0},
}

func TestNewPoint(t *testing.T) {
	for _, val := range testpoints {
		shipPoint := NewPoint(val.row, val.col)
		expectedRow := val.row
		expectedCol := val.col
		row := shipPoint.GetRow()
		col := shipPoint.GetCol()
		if row != expectedRow || col != expectedCol {
			fmt.Println("Result does not match with expected ")
		}
	}
}

func TestGetShipRow(t *testing.T) {
	for _, val := range testpoints {
		shipPoint := NewPoint(val.row, val.col)
		expectedRow := val.row
		resRow := shipPoint.GetRow()
		if resRow != expectedRow {
			fmt.Println("Result does not match with expected ")
		}
	}
}

func TestGetShipCol(t *testing.T) {
	for _, val := range testpoints {
		shipPoint := NewPoint(val.row, val.col)
		expectedCol := val.col
		resCol := shipPoint.GetCol()
		if resCol != expectedCol {
			fmt.Println("Result does not match with expected ")
		}
	}
}
