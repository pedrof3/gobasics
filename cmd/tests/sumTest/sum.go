package tests

func multiply(list ...int) (result int) {
	result = 1
	for _, num := range list {
		result *= num
	}
	return
}
