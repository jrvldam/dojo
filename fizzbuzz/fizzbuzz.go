package fizzbuzz

import (
	"fmt"
	"strings"
)

func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		str := "FizzBuzz"
		if strings.Contains(fmt.Sprintf("%d", n), "3") {
			str = "Fizz" + str
		}
		if strings.Contains(fmt.Sprintf("%d", n), "5") {
			str = str + "Buzz"
		}
		return str
	case n%5 == 0:
		if strings.Contains(fmt.Sprintf("%d", n), "5") {
			return "BuzzBuzz"
		}
		return "Buzz"
	case n%3 == 0:
		if strings.Contains(fmt.Sprintf("%d", n), "3") {
			return "FizzFizz"
		}
		return "Fizz"
	default:
		return fmt.Sprintf("%d", n)
	}
}
