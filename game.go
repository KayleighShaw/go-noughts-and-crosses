package game

import "fmt"

func ResetBoard(board [3][3]string) [3][3]string {
	newBoard := [3][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	return newBoard
}

func TakeTurn(move [1][2]int, player string, board [3][3]string) [3][3]string {
	newBoard := board
	if newBoard[move[0][0]][move[0][1]] != "" {
		fmt.Println("Sorry, that space isn't free. Try again")
	} else {
		newBoard[move[0][0]][move[0][1]] = player
	}

	return newBoard
}
