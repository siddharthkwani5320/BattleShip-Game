package component

import (
	"fmt"
	"testing"
)

type TestCell struct {
	setmark, getmark string
}

var testcell = []TestCell{
	{"X", "X"},
	{"O", "O"},
}

func TestNewCell(t *testing.T) {
	cell := NewCell()
	res := cell.GetMark()
	expect := "-"
	if res != expect {
		fmt.Println("Result:", res, "Expected Result:", expect)
	}
}
func TestSetGetMark(t *testing.T) {
	for _, val := range testcell {
		cell := NewCell()
		cell.SetMark(Mark(val.setmark))
		res := cell.GetMark()
		expect := val.getmark
		if res != expect {
			fmt.Println("Result:", res, "Expected Result:", expect)
		}

	}
}
