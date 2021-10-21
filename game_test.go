package game

import "testing"

func TestResetBoard(t *testing.T) {
	t.Run("empty board is returned from ResetBoard", func(t *testing.T) {
		board := [3][3]string{
			{"", "", ""},
			{"cross", "", ""},
			{"", "", ""},
		}

		got := ResetBoard(board)
		want := [3][3]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		}

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}

func TestTakeTurn(t *testing.T) {
	t.Run("empty board is populated by a cross in the center", func(t *testing.T) {
		board := [3][3]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		}
		move := [1][2]int{
			{1, 1},
		}
		player := "X"

		got := TakeTurn(move, player, board)
		want := [3][3]string{
			{"", "", ""},
			{"", "X", ""},
			{"", "", ""},
		}

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("empty board is populated by a nought in the bottom-center", func(t *testing.T) {
		board := [3][3]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		}
		move := [1][2]int{
			{2, 1},
		}
		player := "O"

		got := TakeTurn(move, player, board)
		want := [3][3]string{
			{"", "", ""},
			{"", "", ""},
			{"", "O", ""},
		}

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}
