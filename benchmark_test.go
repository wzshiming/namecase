package namecase

import (
	"testing"

	nc1 "gopkg.in/wzshiming/namecase.v1"
	nc2 "gopkg.in/wzshiming/namecase.v2"
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

func BenchmarkToSnakeV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testdataBascis {
			nc2.ToLowerSnake(v)
		}
	}
}

func BenchmarkToHumpV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testdataBascis {
			nc2.ToUpperHump(v)
		}
	}
}

func BenchmarkToSnakeV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testdataBascis {
			nc1.Hump2Snake(v)
		}
	}
}

func BenchmarkToHumpV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testdataBascis {
			nc1.Snake2Hump(v)
		}
	}
}
