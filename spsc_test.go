package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewSPSC(t *testing.T) {
	wg := &sync.WaitGroup{}
	prodCon := NewSPSC(8)
	iters := 1000
	wg.Add(2)
	go produce(prodCon, wg, iters)
	go consume(prodCon, wg, iters)
	wg.Wait()
}

func produce(prod Producer, wg *sync.WaitGroup, iters int) {
	for i := 0; i < iters; i++ {
		prod.Produce(i)
	}
	prod.Close()
	wg.Done()
}

func consume(prod Consumer, wg *sync.WaitGroup, iters int) {
	for i := 0; i < iters; i++ {
		if got := prod.Consume(); i != got {
			panic(fmt.Sprintf("invalid data retrieved: expected: %v, got %v", i, got))
		}
	}
	wg.Done()
}

func BenchmarkNewSPSC(b *testing.B) {
	wg := &sync.WaitGroup{}
	prodCon := NewSPSC(8)
	iters := b.N
	wg.Add(2)
	go produce(prodCon, wg, iters)
	go consume(prodCon, wg, iters)
	wg.Wait()
}

func BenchmarkChanSPSC(b *testing.B) {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 8)
	iters := b.N
	wg.Add(2)
	go func() {
		for i := 0; i < iters; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
	go func() {
		for range ch {
		}
		wg.Done()
	}()
	wg.Wait()
}
