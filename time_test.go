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

func TestNow(t *testing.T) {
	ts := Now()
	if ts <= 0 {
		panic("invalid time stamp")
	}
}
