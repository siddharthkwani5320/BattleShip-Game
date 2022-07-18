package component

import (
	"fmt"
	"testing"
)

type player struct {
	name     string
	noOfHit  int
	noOfMiss int
}

var players = []player{
	{"Siddharth", 0, 0},
	{"Vincenzo", 0, 0},
	{"Jang Joo", 0, 0},
	{"Sung Joong Ki", 0, 0},
}

func TestNewPlayer(t *testing.T) {
	for _, val := range players {
		newPlayer := NewPlayer(val.name)
		expectName := val.name
		resName := newPlayer.GetName()
		expectHit := val.noOfHit
		resNoOfHit := newPlayer.GetHit()
		expectMiss := val.noOfHit
		resNoOfMiss := newPlayer.GetMiss()
		if resName != string(expectName) {
			fmt.Println("Result:", resName, "Expected Result:", expectName)
		}
		if resNoOfHit != expectHit {
			fmt.Println("Result:", resNoOfHit, "Expected Result:", expectHit)
		}
		if resNoOfMiss != expectMiss {
			fmt.Println("Result:", resNoOfMiss, "Expected Result:", expectMiss)
		}
	}
}

func TestIncrementHit(t *testing.T) {
	for _, val := range players {
		newPlayer := NewPlayer(val.name)
		expectedHit := 3
		for i := 0; i < 3; i++ {
			newPlayer.IncrementHit()
		}
		if newPlayer.GetHit() != expectedHit {
			fmt.Println("Result:", newPlayer.GetHit(), "Expected Result:", expectedHit)
		}
	}
}

func TestIncrementMiss(t *testing.T) {
	for _, val := range players {
		newPlayer := NewPlayer(val.name)
		expectedMiss := 3
		for i := 0; i < 3; i++ {
			newPlayer.IncrementMiss()
		}
		if newPlayer.GetMiss() != expectedMiss {
			fmt.Println("Result:", newPlayer.GetHit(), "Expected Result:", expectedMiss)
		}
	}
}
