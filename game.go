package main

import "fmt"

type Board struct {
	P1 string
	P2 string
	State [3][3]string
}

func getInitialBoard() [3][3]string {
	board := [3][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	return board
}

func ResetBoard(board [3][3]string) [3][3]string {
	newBoard := [3][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	return newBoard
}

/*
// mutation approach
board = NewBoard()
TakeTurn(player, move, board)
TakeTurn(player, move, board) // TakeTurn is a function that modifies board directly

// immutable data approach <- this
board = NewBoard()
board = TakeTurn(player, move, board)
board = TakeTurn(player, move, board)  
    // TakeTurn is a function that takes old board, and returns a brand new board with the changes applied
*/

func TakeTurn(move [1][2]int, player string, board [3][3]string) [3][3]string {
	//copy old board to new board
	//set value
	//return new board
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

func main() {
	//TODO: randomise noughts and crosses for players

	//Taking input from player 1
	fmt.Println("Enter a name for Player 1:")
	var p1 string
	fmt.Scanln(&p1)
	fmt.Println("Hi, " + p1 + ". You are Crosses!")

	//Taking input for player 2
	fmt.Println("Enter a name for Player 2:")
	var p2 string
	fmt.Scanln(&p2)
	fmt.Println("Hi, " + p2 + ". You are Crosses!")

	// creating an instance for initialState of the game
	initialState := Board{p1, p2, getInitialBoard()}

	// this is a pointer to the Board struct
	pts := &initialState
	fmt.Println(pts)

	// this uses the pointer to accessing the struct fields (not using dereferencing explicitly)
	fmt.Println(pts.P1)

	// explicitly using dereferencing concept (the result will be the same)
	fmt.Println((*pts).P1)

	// can also modify the value of P1
	// TODO use this for board state changes
	pts.P1 = "Bob"
	fmt.Println(pts.P1)
}