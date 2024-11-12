package main

import (
	"fmt"
	"regexp"
)

func RockPaperScissors() {
	var player1Str, player2Str string
	fmt.Println("Player 1 select rock, paper, or scissors:")
	fmt.Scanln(&player1Str)
	fmt.Println("Player 2, select rock, paper, or scissors:")
	fmt.Scanln(&player2Str)
	player1 := conversion(player1Str)
	player2 := conversion(player2Str)
	comparision := Comparision(player1, player2)
	fmt.Println(comparision)
}

func conversion(playerStr string) int {
	if regexp.MustCompile(`(?i)^r..k$`).MatchString(playerStr) {
		return 1
	}
	if regexp.MustCompile(`(?i)^p.*r$`).MatchString(playerStr) {
		return 2
	}
	if regexp.MustCompile(`(?i)^s.*r$`).MatchString(playerStr) {
		return 3
	}
	switch playerStr {
	case "rock":
		return 1
	case "paper":
		return 2
	case "scissors", "scissor":
		return 3
	default:
		return 0
	}
}

/*func Comparision(player1 int, player2 int) string {
	if player1 == player2 {
		return "it's a tie"
	} else if player1 == 1 && player2 == 3 {
		return "Player 1 wins"
	} else if player1 == 3 && player2 == 1 {
		return "Player 2 wins"
	} else if player1 > player2 {
		return "Player 1 wins"
	} else if player2 > player1 {
		return "Player 2 wins"
	} else {
		return "Invalid input"
	}
}*/

func Comparision(player1 int, player2 int) string {
	if player1 == player2 {
		return "It's a tie"
	}

	winMap := map[int]int{
		1: 3, // Rock beats Scissors
		2: 1, // Paper beats Rock
		3: 2, // Scissors beats Paper
	}

	if winMap[player1] == player2 {
		return "Player 1 wins"
	} else if winMap[player2] == player1 {
		return "Player 2 wins"
	} else {
		return "Invalid input"
	}
}

/*func main() {
	RockPaperScissors()
}*/
