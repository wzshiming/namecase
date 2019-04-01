package namecase

import (
	"testing"
)

func BenchmarkToSnake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testdataBascis {
			ToLowerSnake(v)
		}
	}
}

func BenchmarkToHump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testdataBascis {
			ToUpperHump(v)
		}
	}
}
