package user

import "fmt"

func GetInput() []int {
	var column, row int
	for {
		fmt.Print("Enter row: ")
		fmt.Scanln(&row)
		if row > 0 && row <= 3 {
			break
		} else {
			fmt.Println("Choose row 1, 2 or 3")
		}
	}
	for {
		fmt.Print("Enter column: ")
		fmt.Scanln(&column)
		if column > 0 && column <= 3 {
			break
		} else {
			fmt.Println("Choose column 1, 2 or 3")
		}
	}
	return []int{row - 1, column - 1}
}
