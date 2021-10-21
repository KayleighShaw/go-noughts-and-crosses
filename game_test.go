package game

import "testing"

func TestBoard(t *testing.T) {
	t.Run("empty board is returned from ResetBoard", func(t *testing.T) {
		board := [4][3]string{
			{"", "", ""},
			{"cross", "", ""},
			{"", "", ""},
			{"", "", ""},
		}

		got := ResetBoard(board)
		want := [4][3]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		}

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}
