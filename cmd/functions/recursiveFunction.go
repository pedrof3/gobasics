package functions

func Fatorial(number int) int {
	if number <= 1 {
		return 1
	}
	return Fatorial(number-1) * number
}
