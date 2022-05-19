package board

import (
	"fmt"
	"strings"
)

func DrawBoard(board [][3]string) {
	for _, row := range board {
		for y, s := range row {
			if s == "" {
				row[y] = "_"
			}
		}
		fmt.Println(strings.Join(row[:], ""))
	}
}

func CheckValidMove(move []int, board [][3]string) bool {
	xCoord := move[0]
	yCoord := move[1]

	if board[xCoord][yCoord] != "" {
		return false
	}
	return true
}
