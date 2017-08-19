package selectrand

import (
	"math/rand"
	"testing"
)

func BenchmarkSelectRand(b *testing.B) {
	s := Source{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Uint64()
	}
}

func BenchmarkMathRand(b *testing.B) {
	s := rand.NewSource(0).(rand.Source64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Uint64()
	}
}
