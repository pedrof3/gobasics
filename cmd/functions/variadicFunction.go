package functions

func SumSequence(list ...int) (total int) {
	for _, number := range list {
		total += number
	}
	return
}
