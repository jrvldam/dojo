package leapyear

type Year struct {
	year int
}

func (y Year) IsLeapYear() bool {
	return y.divisibleBy(4) && !(y.divisibleBy(100) && y.notDivisibleBy(400))
}

func (y Year) divisibleBy(number int) bool {
	return y.year%number == 0
}

func (y Year) notDivisibleBy(number int) bool {
	return y.year%number != 0
}

func NewYear(year int) Year {
	return Year{
		year,
	}
}
