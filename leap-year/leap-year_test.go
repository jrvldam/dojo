package leapyear_test

import (
	"fmt"
	"testing"

	leapyear "github.com/jrvldam/dojo/leap-year"
)

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year       int
		isLeapYear bool
	}{
		{year: 1997, isLeapYear: false},
		{year: 1996, isLeapYear: true},
		{year: 1600, isLeapYear: true},
		{year: 1800, isLeapYear: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("year %d", test.year), func(t *testing.T) {
			year := leapyear.NewYear(test.year)
			got := year.IsLeapYear()

			if got != test.isLeapYear {
				t.Errorf("want %v, got %v", test.isLeapYear, got)
			}
		})
	}
}
