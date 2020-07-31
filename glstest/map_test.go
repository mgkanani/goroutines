package glstest

import (
	"flag"
	"sync"
	"testing"
)

var (
	mp        *sync.Map
	stdMap    map[int]int
	entries   = flag.Int("entries", 1, "number of expunged entries in sync.map before benchmarking")
	isSyncMap = flag.Bool("sync", true, "is sync.map or std map")
)

func Init() {
	mp = &sync.Map{}
	stdMap = make(map[int]int)

	flag.Parse()

	st := 100000
	end := st + *entries

	if *isSyncMap {
		mpInit(st, end)
	} else {
		stdMapInit(st, end)
	}
}

func mpInit(st, end int) {
	for i := st; i <= end; i++ {
		mp.Store(i, i)
	}

	for i := st; i <= end; i++ {
		mp.Load(i)
	}

	for i := st; i <= end; i++ {
		mp.Delete(i) // even though entries are to be deleted, sync-map will be containing all the entries
	}
}

func stdMapInit(st, end int) {
	for i := st; i <= end; i++ {
		stdMap[i] = i
	}
	for i := st; i <= end; i++ {
		delete(stdMap, i)
	}
}

func BenchmarkSyncMapLoad(b *testing.B) {
	Init()
	for i := 0; i < b.N; i++ {
		mp.Load(i)
	}
}

func BenchmarkStdMapLoad(b *testing.B) {
	Init()
	for i := 0; i < b.N; i++ {
		_ = stdMap[i]
	}
}
