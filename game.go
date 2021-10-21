package game

func ResetBoard(board [4][3]string) [4][3]string {
	newBoard := [4][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	return newBoard
}
