package runtime

import (
	"testing"
)

func TestGoRoutine(t *testing.T) {
	g := goRoutine()
	if g == nil {
		panic("invalid go-routine pointer")
	}
}

func BenchmarkCustom(b *testing.B) {
	N := b.N
	var goRtn *g
	mainRtn := goRoutine()
	lck := &mutex{}

	go func(max int) {
		goRtn = goRoutine()
		for i := 0; i < max; i++ {
			lock(lck)
			goready(mainRtn, 0)
			gopark(lck, waitReasonChanReceive, traceEvGoBlockRecv, 0)
		}
	}(N)

	lock(lck)
	gopark(lck, waitReasonChanReceive, traceEvGoBlockRecv, 0)
	for i := 0; i < N-1; i++ {
		lock(lck)
		goready(goRtn, 0)
		gopark(lck, waitReasonChanReceive, traceEvGoBlockRecv, 0)
	}
}

func BenchmarkChan(b *testing.B) {
	// If just one channel is used then it won't be fair comparison.
	// In std-channel based communication, sender underneath calls goready.
	// while receiver calls gopark.
	child := make(chan int)
	parent := make(chan int)
	N := b.N
	go func(blockSelf <-chan int, awakeParent chan<- int) {
		for i := 0; i < N; i++ {
			<-blockSelf
			awakeParent <- i
		}
		close(parent)
	}(child, parent)
	for i := 0; i < N; i++ {
		child <- i
		<-parent
	}
	close(child)
}
