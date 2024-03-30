package bench

func normalFatorial(num int) int {
	result := 1
	result = 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result
}

func recursiveFatorial(num int) int {
	if num <= 1 {
		return 1
	}
	return recursiveFatorial(num-1) * num
}
