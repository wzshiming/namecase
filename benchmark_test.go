package namecase

import (
	"testing"

	nc1 "gopkg.in/wzshiming/namecase.v1"
)

func BenchmarkToSnake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range data1 {
			ToLowerSnake(v)
		}
	}
}

func BenchmarkToHump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range data1 {
			ToUpperHump(v)
		}
	}
}

func BenchmarkToSnakeV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range data1 {
			nc1.Hump2Snake(v)
		}
	}
}

func BenchmarkToHumpV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range data1 {
			nc1.Snake2Hump(v)
		}
	}
}
