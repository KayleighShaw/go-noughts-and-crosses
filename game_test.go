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
		player := "cross"

		got := TakeTurn(move, player, board)
		want := [3][3]string{
			{"", "", ""},
			{"", "cross", ""},
			{"", "", ""},
		}

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}
