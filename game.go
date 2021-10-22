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

func rowWin(board [3][3]string) string {
	for row := 0; row < 3; row++ {
		if board[row][0] == board[row][1] && board[row][1] == board[row][2] {
			if board[row][0] == "X" {
				fmt.Println("Crosses wins!")
				return "Crosses wins!"
			} else {
				fmt.Println("Noughts wins!")
				return "Noughts wins!"
			}
		}
	}
	return isItADraw(board) // perhaps make this a tuple? Return a boolean (win) and a string?
}

func columnWin(board [3][3]string) string {
	for column := 0; column < 3; column++ {
		if board[0][column] == board[1][column] && board[1][column] == board[2][column] {
			if board[0][column] == "X" {
				fmt.Println("Crosses wins!")
				return "Crosses wins!"
			} else {
				fmt.Println("Noughts wins!")
				return "Noughts wins!"
			}
		}
	}
	return isItADraw(board)
}

func isItADraw(board [3][3]string) string {
	gameContinue := false
	for row := 0; row < 3; row++ {
		if board[row][0] == "" || board[row][1] == "" || board[row][2] == "" {
			gameContinue = true
		}
	}
	if gameContinue {
		return "Next player"
	} else if gameContinue == false {
		return "It's a draw!"
	}
	return "This isn't working"
}
