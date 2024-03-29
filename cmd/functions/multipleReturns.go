package functions

func SchoolAverage(n1, n2 float64) (float64, bool) {
	average := (n1 + n2) / 2
	if average < 6 {
		return average, false
	}
	return average, true
}
