package guessnumber

import (
	"math/rand"
	"time"
)

const (
	Wins   = "wins"
	Loses  = "loses"
	Lower  = "lower"
	Grater = "grater"
)

type game struct {
	randomNumber int
	attempts     int
}

func (g *game) GuessNumber(number int) string {
	var result string

	g.attempts = g.attempts - 1

	switch {
	case g.randomNumber < number:
		result = Lower
	case g.randomNumber > number:
		result = Grater
	default:
		result = Wins
	}

	if g.attempts == 0 && result != Wins {
		result = Loses
	}

	return result
}

func NewGame(numberGenerator func() int) game {
	return game{
		randomNumber: numberGenerator(),
		attempts:     3,
	}
}

func GenerateNumber() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(10-1) + 1
}
