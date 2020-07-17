package runtime

import (
	"testing"
	"time"
)

func BenchmarkCustomNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Now()
	}
}

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}
