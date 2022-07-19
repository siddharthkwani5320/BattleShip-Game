package service

import (
	"battlegame/component"
	resultanalyser "battlegame/resultAnalyser"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NewGame() {
	player := CreatePlayer()
	board := CreateBoard()

	Play(player, board)
}

func CreatePlayer() *component.Player {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name:")
	name, err := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")
	if err != nil {
		fmt.Println("Error:", err)
	}

	player := component.NewPlayer(name)

	return player
}

func Play(player *component.Player, board *component.Board) {
	reader := bufio.NewReader(os.Stdin)
	var i, tries int
	tries = (board.GetRow() * board.GetCol()) / 2
	if tries < 15 {
		tries = 15
	}
	for i = 1; i <= tries; i++ {
		fmt.Println("No. Of Tries Left With you:", (tries - i + 1))
		var row, col int
		var err error
		fmt.Println("Current Board")
		board.PrintPlayerBoard()
		for {
			fmt.Println("Enter the Row you want to hit:")
			//User Input of Row
			Brow, _ := reader.ReadString('\n')
			Brow = strings.TrimSpace(Brow)
			row, err = strconv.Atoi(Brow)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			//Check If its a Valid Row
			if row >= board.GetRow() || row < 0 {
				fmt.Println("Please Enter within 0 to ", board.GetRow())
				continue
			}
			break
		}
		for {
			//User Input of Column
			fmt.Println("Enter the Column you want to hit:")
			Bcol, _ := reader.ReadString('\n')
			Bcol = strings.TrimSpace(Bcol)
			col, err = strconv.Atoi(Bcol)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			//Check If its a Valid Column
			if int(col) >= board.GetCol() || col < 0 {
				fmt.Println("Please Enter within 0 to ", board.GetCol())
				continue
			}
			break
		}

		mark := board.Cell[row][col].GetMark() //If that Cell Contains a ship or not
		if mark == string(component.HitMark) || mark == string(component.MissMark) {
			fmt.Println("This Cell is already hitted.")
			continue
		}
		if mark != string(component.NoMark) {
			fmt.Println("Yayy!!...Its A Hit.")
			board.Cell[row][col].SetMark(component.HitMark)
			player.IncrementHit()
			check := resultanalyser.Analysis(board, component.ShipPoint)
			if check {
				name := player.GetName()
				fmt.Println(name)
				fmt.Println(player.GetName(), " Wins the Game ")
				fmt.Println("Total Hits:", player.GetHit())
				fmt.Println("Total Miss:", player.GetMiss())
				break
			}
			continue
		}
		fmt.Println("Oopps!!!..Its a Miss.")
		player.IncrementMiss()
		board.Cell[row][col].SetMark(component.MissMark)
	}
	fmt.Println("You Lost the Game...Number of tries Ends")
	fmt.Println("Total Hits:", player.GetHit())
	fmt.Println("Total Miss:", player.GetMiss())
}
