package tests

import (
	"log"
	"testing"
)

type validator struct {
	input  []int
	output int
}

func TestMultiply(t *testing.T) {
	testCases := []validator{
		{[]int{1, 2, 3}, 6},
		{[]int{2, 2, 2, 2}, 16},
		{[]int{0, 1, 3, 5, 7}, 30},
	}

	for _, testCase := range testCases {
		if multiply(testCase.input...) != testCase.output {
			log.Fatal("erro encontrado")
		}
	}
}
