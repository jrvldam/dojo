package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/jrvldam/dojo/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		number int
		want   string
	}{
		{number: 1, want: "1"},
		{number: 3, want: "FizzFizz"},
		{number: 5, want: "BuzzBuzz"},
		{number: 6, want: "Fizz"},
		{number: 10, want: "Buzz"},
		{number: 15, want: "FizzBuzzBuzz"},
		{number: 16, want: "16"},
		{number: 18, want: "Fizz"},
		{number: 25, want: "BuzzBuzz"},
		{number: 30, want: "FizzFizzBuzz"},
		{number: 135, want: "FizzFizzBuzzBuzz"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("testing for %d", test.number), func(t *testing.T) {
			got := fizzbuzz.FizzBuzz(test.number)

			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}

		})
	}
}
