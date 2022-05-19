package main

import (
	"fmt"

	"ticTacGo/board"
	"ticTacGo/game"
	"ticTacGo/user"
)

func main() {
	gameBoard := make([][3]string, 3)
	playersTurn := "X"

	for {
		playersTurn = game.GetNextPlayersTurn(playersTurn)
		var move []int
		for {
			move = user.GetInput()
			if valid := board.CheckValidMove(move, gameBoard); valid {
				break
			}
			fmt.Println("Space already taken")
		}
		gameBoard[move[0]][move[1]] = playersTurn
		board.DrawBoard(gameBoard)
		if win, winner := game.CheckWinner(gameBoard); win {
			fmt.Printf("Winner is %s!!!", winner)
			break
		}
	}
}
