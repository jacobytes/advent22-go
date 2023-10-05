package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	content, error := os.ReadFile("input.txt")

	if error != nil {
		log.Fatal(error)
	}

	games := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	moves := make(map[string]string)

	moves["A"] = "rock"
	moves["B"] = "paper"
	moves["C"] = "scissors"
	moves["X"] = "rock"
	moves["Y"] = "paper"
	moves["Z"] = "scissors"

	moveScores := make(map[string]int)

	moveScores["rock"] = 1
	moveScores["paper"] = 2
	moveScores["scissors"] = 3

	outcomes := make(map[string]string)

	outcomes["X"] = "loss"
	outcomes["Y"] = "draw"
	outcomes["Z"] = "win"

	winningMoves := make(map[string]string)

	winningMoves["rock"] = "paper"
	winningMoves["paper"] = "scissors"
	winningMoves["scissors"] = "rock"

	losingMoves := make(map[string]string)

	losingMoves["rock"] = "scissors"
	losingMoves["paper"] = "rock"
	losingMoves["scissors"] = "paper"

	var totalScore = 0

	for _, input := range games {
		var gameScore = 0
		playerInput := strings.Fields(input)

		opponentChoice := moves[playerInput[0]]
		outcome := outcomes[playerInput[1]]

		playerChoice := ""

		if outcome == "draw" {
			playerChoice = opponentChoice
		}

		if outcome == "win" {
			playerChoice = winningMoves[opponentChoice]
		}

		if outcome == "loss" {
			playerChoice = losingMoves[opponentChoice]
		}

		choiceScore := moveScores[playerChoice]

		game := opponentChoice + " " + playerChoice

		if opponentChoice == playerChoice {
			gameScore = 3
		} else if game == "rock paper" || game == "paper scissors" || game == "scissors rock" {
			gameScore = 6
		}

		totalScore += gameScore + choiceScore
	}

	fmt.Println(totalScore)
}
