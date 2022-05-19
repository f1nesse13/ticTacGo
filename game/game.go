package game

// [[0,0],[0,1][0,2]], [[1,0],[1,1][1,2]],[[2,0],[2,1][2,2]], <-- ROWs
// [[0,0],[1,0][2,0]], [[0,1],[1,1][2,1]],[[0,2],[1,2][2,2]] <-- COLs
// [[0,0],[1,1][2,2]], [[0,2],[1,1][2,0]] <-- DIAGs

const X = "X"
const O = "O"

var winningPositions = [][3][2]int{{{0, 0}, {0, 1}, {0, 2}}, {{1, 0}, {1, 1}, {1, 2}}, {{2, 0}, {2, 1}, {2, 2}}, {{0, 0}, {1, 0}, {2, 0}}, {{0, 1}, {1, 1}, {2, 1}}, {{0, 2}, {1, 2}, {2, 2}}, {{0, 0}, {1, 1}, {2, 2}}, {{0, 2}, {1, 1}, {2, 0}}}

func CheckWinner(board [][3]string) (bool, string) {
	for _, winPattern := range winningPositions {
		xCount, oCount := 0, 0
		for _, pos := range winPattern {
			posY, posX := pos[0], pos[1]

			if board[posY][posX] == "X" {
				xCount += 1
			} else if board[posY][posX] == "O" {
				oCount += 1
			}
		}
		if xCount == 3 {
			return true, "X"
		} else if oCount == 3 {
			return true, "O"
		}
	}
	return false, ""
}

func GetNextPlayersTurn(s string) string {
	if s == "O" {
		return "X"
	}
	return "O"

}
