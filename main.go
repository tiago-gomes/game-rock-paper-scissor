package main

import "fmt"

func getPlayerName(prompt string) string {
	var name string
	fmt.Print(prompt)
	fmt.Scanln(&name)
	return name
}

func getPlayerMove(prompt string) string {
	var move string

	for {
		fmt.Print(prompt)
		fmt.Scanln(&move)

		if move == "rock" || move == "paper" || move == "scissors" {
			return move // Return the valid move and exit the function
		}

		fmt.Println("Invalid move. Please enter rock, paper, or scissors.")
	}
}

func checkPlay(
	playerOneMove string,
	playerTwoMove string,
	playerOneName string,
	playerTwoName string,
) {
	if playerOneMove == playerTwoMove {
		fmt.Println("It's a tie!")
		return
	}

	// player one wins
	if playerOneMove == "rock" && playerTwoMove == "scissors" ||
		playerOneMove == "scissors" && playerTwoMove == "paper" ||
		playerOneMove == "paper" && playerTwoMove == "rock" {
		fmt.Println(playerOneName, " wins!")
		return
	}

	// player two wins
	fmt.Println(playerTwoName, " wins!")
	return
}

func main() {

	// Welcome screen
	fmt.Println("Welcome to Rock, Paper, Scissor Game")

	// Get player names
	playerOneName := getPlayerName("Enter player 1 name: ")
	playerTwoName := getPlayerName("Enter player 2 name: ")

	// Get players move
	playerOneMove := getPlayerMove(fmt.Sprintf("Enter %s move: ", playerOneName))
	playerTwoMove := getPlayerMove(fmt.Sprintf("Enter %s move: ", playerTwoName))

	// Decide win or tie
	checkPlay(
		playerOneMove,
		playerTwoMove,
		playerOneName,
		playerTwoName,
	)
}
