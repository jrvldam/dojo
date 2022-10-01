package guessnumber_test

import (
	"testing"

	guessnumber "github.com/jrvldam/dojo/guess-random-number"
)

func makeGenerateNumberMock(n int) func() int {
	return func() int {
		return n
	}
}

func TestGuessNumber(t *testing.T) {
	t.Run("player wins at first attempt", func(t *testing.T) {
		mock := makeGenerateNumberMock(5)
		game := guessnumber.NewGame(mock)

		got := game.GuessNumber(5)

		if got != guessnumber.Wins {
			t.Errorf("want %q, got %q", guessnumber.Wins, got)
		}
	})

	t.Run("player guess is grater than random number", func(t *testing.T) {
		mock := makeGenerateNumberMock(1)
		game := guessnumber.NewGame(mock)

		got := game.GuessNumber(5)

		if got != guessnumber.Lower {
			t.Errorf("want %q, got %q", guessnumber.Lower, got)
		}
	})

	t.Run("player guess is lower than random number", func(t *testing.T) {
		mock := makeGenerateNumberMock(5)
		game := guessnumber.NewGame(mock)

		got := game.GuessNumber(1)

		if got != guessnumber.Grater {
			t.Errorf("want %q, got %q", guessnumber.Grater, got)
		}
	})

	t.Run("player wins on second attempt", func(t *testing.T) {
		mock := makeGenerateNumberMock(5)
		game := guessnumber.NewGame(mock)

		attempt := game.GuessNumber(1)

		if attempt != guessnumber.Grater {
			t.Errorf("want %q, got %q", guessnumber.Grater, attempt)
		}

		result := game.GuessNumber(5)

		if result != guessnumber.Wins {
			t.Errorf("want %q, got %q", guessnumber.Wins, result)
		}
	})

	t.Run("player wins on third attempt", func(t *testing.T) {
		mock := makeGenerateNumberMock(5)
		game := guessnumber.NewGame(mock)

		attemptOne := game.GuessNumber(1)

		if attemptOne != guessnumber.Grater {
			t.Errorf("want %q, got %q", guessnumber.Grater, attemptOne)
		}

		attemptTwo := game.GuessNumber(7)

		if attemptTwo != guessnumber.Lower {
			t.Errorf("want %q, got %q", guessnumber.Lower, attemptTwo)
		}

		attemptThree := game.GuessNumber(5)

		if attemptThree != guessnumber.Wins {
			t.Errorf("want %q, got %q", guessnumber.Wins, attemptThree)
		}
	})

	t.Run("player loses", func(t *testing.T) {
		mock := makeGenerateNumberMock(5)
		game := guessnumber.NewGame(mock)

		attemptOne := game.GuessNumber(1)

		if attemptOne != guessnumber.Grater {
			t.Errorf("want %q, got %q", guessnumber.Grater, attemptOne)
		}

		attemptTwo := game.GuessNumber(7)

		if attemptTwo != guessnumber.Lower {
			t.Errorf("want %q, got %q", guessnumber.Lower, attemptTwo)
		}

		attemptThree := game.GuessNumber(6)

		if attemptThree != guessnumber.Loses {
			t.Errorf("want %q, got %q", guessnumber.Loses, attemptThree)
		}
	})
}
