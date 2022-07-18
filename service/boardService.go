package service

import (
	"battlegame/component"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CreateBoard() *component.Board {
	reader := bufio.NewReader(os.Stdin)
	var row, col int
	var err error
	for {
		//User Input of Row
		fmt.Println("Enter the Rows of Board:")
		Brow, _ := reader.ReadString('\n')
		Brow = strings.TrimSpace(Brow)
		row, err = strconv.Atoi(Brow)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if row < 5 {
			fmt.Println("Error:Row Must be greater than 4")
			continue
		}
		break
	}
	for {
		//User Input of Column
		fmt.Println("Enter the Column of Board:")
		Bcol, _ := reader.ReadString('\n')
		Bcol = strings.TrimSpace(Bcol)
		col, err = strconv.Atoi(Bcol)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if col < 5 {
			fmt.Println("Error:Column Must be greater than 4")
			continue
		}
		break
	}

	board, err := component.NewBoard(row, col)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("Original Board")
	board.PrintBoard()
	point := CreateShip(board, 5, "S1")
	component.ShipPoint[5] = point
	point = CreateShip(board, 4, "S2")
	component.ShipPoint[4] = point
	point = CreateShip(board, 3, "S3")
	component.ShipPoint[3] = point
	point = CreateShip(board, 2, "S4")
	component.ShipPoint[2] = point
	point = CreateShip(board, 1, "S5")
	component.ShipPoint[1] = point
	fmt.Println("Ship Inserted Board")
	board.PrintBoard()
	return board
}
