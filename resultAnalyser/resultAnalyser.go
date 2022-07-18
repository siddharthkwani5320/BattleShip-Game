package resultanalyser

import (
	"battlegame/component"
	"fmt"
	"strconv"
)

func Analysis(board *component.Board, ShipPoint map[int][]*component.Point) bool {
	cnt := 0
	for key, points := range ShipPoint {
		cnt = 0
		for _, point := range points {
			mark := board.GetCell(point.GetRow(), point.GetCol()).GetMark()
			if mark == string(component.HitMark) {
				cnt++
			}
		}
		if key == cnt {
			delete(ShipPoint, key)
			fmt.Println("Ship " + strconv.Itoa(key) + " is destroyed")
			return false
		}
	}
	if len(ShipPoint) == 0 {
		return true
	}
	return false
}
