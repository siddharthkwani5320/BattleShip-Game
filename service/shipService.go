package service

import (
	"battlegame/component"
	"fmt"
	"math/rand"
	"time"
)

func CreateShip(board *component.Board, size int, shipMark string) []*component.Point {
	var shipPoints []*component.Point
	for {
		rand.Seed(time.Now().UnixNano())
		//Random Ship Location
		row := rand.Intn(board.GetRow())
		col := rand.Intn(board.GetCol())

		randomSelect := rand.Intn(2)

		if randomSelect == 0 {
			//Check Row
			start := Up(board, row, col)
			stop := Down(board, row, col)

			if start == -1 || stop == -1 {
				continue
			}
			if (start + size - 1) <= stop {
				fmt.Println("Row:", row, col, start, stop)
				shipPoints = InsertShipRow(board, start, start+size, col, shipMark)
				board.PrintBoard()

				break
			}
		} else {

			//Check Column
			start := Left(board, row, col)
			stop := Right(board, row, col)

			if start == -1 || stop == -1 {
				continue
			}

			if (start + size - 1) <= stop {
				fmt.Println("Column:", row, col, start, stop)
				shipPoints = InsertShipColumn(board, start, start+size, row, shipMark)
				board.PrintBoard()
				break
			}
		}
	}
	return shipPoints
}

//Inserting Ship Into Row of the Board
func InsertShipRow(board *component.Board, start, stop, col int, shipMark string) []*component.Point {
	var tempPoints []*component.Point
	for i := start; i < stop; i++ {
		temp := component.NewPoint(i, col)
		tempPoints = append(tempPoints, temp)
		board.Cell[i][col].SetMark(component.Mark(shipMark))
	}
	return tempPoints
}

//Inserting Ship Into Column of the Board
func InsertShipColumn(board *component.Board, start, stop, row int, shipMark string) []*component.Point {
	var tempPoints []*component.Point
	for i := start; i < stop; i++ {
		temp := component.NewPoint(row, i)
		tempPoints = append(tempPoints, temp)
		board.Cell[row][i].SetMark(component.Mark(shipMark))
	}
	return tempPoints
}

func Right(board *component.Board, row, col int) int {
	pos := -1
	for i := col; i < board.GetCol(); i++ {
		if board.Cell[row][i].GetMark() == string(component.NoMark) {
			pos = i
			continue
		}
		break
	}
	return pos
}

func Left(board *component.Board, row, col int) int {
	pos := -1
	for i := col; i >= 0; i-- {
		if board.Cell[row][i].GetMark() == string(component.NoMark) {
			pos = i
			continue
		}
		break
	}
	return pos
}

func Down(board *component.Board, row, col int) int {
	pos := -1
	for i := row; i < board.GetRow(); i++ {
		if board.Cell[i][col].GetMark() == string(component.NoMark) {
			pos = i
			continue
		}
		break
	}
	return pos
}

func Up(board *component.Board, row, col int) int {
	pos := -1
	for i := row; i >= 0; i-- {
		if board.Cell[i][col].GetMark() == string(component.NoMark) {
			pos = i
			continue
		}
		break
	}
	return pos
}
