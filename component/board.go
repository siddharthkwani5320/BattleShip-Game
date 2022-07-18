package component

import "fmt"

type Board struct {
	bRow int
	bCol int
	Cell [][]*Cells
}

func NewBoard(row, col int) (*Board, error) {
	var board Board
	var i, j int
	board.bRow = row
	board.bCol = col
	for i = 0; i < board.bRow; i++ {
		var tempRow []*Cells
		for j = 0; j < board.bCol; j++ {
			cell := NewCell()
			tempRow = append(tempRow, cell)
		}
		board.Cell = append(board.Cell, tempRow)
	}
	return &board, nil
}

func (b *Board) GetCell(row, col int) *Cells {
	return b.Cell[row][col]
}
func (b *Board) GetRow() int {
	return b.bRow
}
func (b *Board) GetCol() int {
	return b.bCol
}

func (b *Board) PrintBoard() {
	var i, j int
	for i = 0; i < b.bRow; i++ {
		for j = 0; j < b.bCol; j++ {
			m := b.Cell[i][j].GetMark()
			if m == string(NoMark) {
				fmt.Print(m, "  ")
				continue
			}
			fmt.Print(m, " ")
		}
		fmt.Println()
	}
}

func (b *Board) PrintPlayerBoard() {
	var i, j int
	for i = 0; i < b.bRow; i++ {
		for j = 0; j < b.bCol; j++ {
			m := b.Cell[i][j].GetMark()
			if m == string(HitMark) || m == string(MissMark) {
				fmt.Print(m, "  ")
			} else {
				fmt.Print("-  ")
			}
		}
		fmt.Println()
	}
}
