package main

import (
	"math/rand"
	"testing"
)

func TestRockPaperScissors1(t *testing.T) {
	player1 := 1
	player2 := 3
	comparision := Comparision(player1, player2)
	if comparision != "Player 1 wins" {
		t.Errorf("Expected Player 1 wins but got %s", comparision)
	}
}

func TestRockPaperScissors2(t *testing.T) {
	player1 := rand.Intn(3)
	player2 := rand.Intn(3)
	comparision := Comparision(player1, player2)
	if comparision != "Player 2 wins" && comparision != "Player 1 wins" && comparision != "It's a tie" {
		t.Errorf("Unexpected result: %s", comparision)
	}
}
