package src

import (
	"testing"
)

var list []int64

func init() {
	for i := 0; i < Length; i++ {
		list = append(list, int64(i))
	}
}

func BenchmarkConcurencyWithHTTP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ConcurencyHTTP(list, 10)
	}
}

func BenchmarkNotConcurencyHTTP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NotConcurencyHTTP(list, 10)
	}
}
