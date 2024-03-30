package bench

import (
	"testing"
)

func BenchmarkNormalFatorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalFatorial(20)
	}
}

func BenchmarkRecursiveFatorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursiveFatorial(20)
	}
}
