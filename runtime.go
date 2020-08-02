package goroutines

import (
	"unsafe"
)

// Event types in the trace, args are given in square brackets.
// go tool trace will behave as per data-values inside square brackets.
const (
	waitReasonChanReceive = 14
	waitReasonChanSend    = 15
	traceEvGoBlockSend    = 22 // goroutine blocks on chan send [timestamp, stack]
	traceEvGoBlockRecv    = 23 // goroutine blocks on chan recv [timestamp, stack]

	waiting = 4 // must be in sync with _Gwaiting in runtime/runtime2.go
)

type g struct{}

type mutex struct {
	key uintptr
}

//go:linkname lock runtime.lock
func lock(m *mutex)

//go:linkname unlock runtime.unlock
func unlock(m *mutex)

//go:linkname goparkunlock runtime.goparkunlock
func goparkunlock(lock *mutex, reason uint8, traceEv byte, traceskip int)

//go:linkname gopark runtime.gopark
func gopark(unlockf func(*g, unsafe.Pointer) bool, lock unsafe.Pointer, reason uint8, traceEv int, traceSkip int)

//go:linkname goready runtime.goready
func goready(gp *g, traceskip int)

//go:linkname readgstatus runtime.readgstatus
func readgstatus(gp *g) uint32

func Park(isSender bool, unlockFunc func(*g, unsafe.Pointer) bool, lck *mutex) {
	lock(lck)
	//clouse := (func(*g, unsafe.Pointer) bool)(unlockFunc) // todo: make it exported by making *g unsafe.Pointer
	if isSender {
		gopark(unlockFunc, unsafe.Pointer(lck), waitReasonChanSend, traceEvGoBlockSend, 1)
	} else {
		gopark(unlockFunc, unsafe.Pointer(lck), waitReasonChanReceive, traceEvGoBlockRecv, 1)
	}
}

func Awake(rtn unsafe.Pointer) {
	gp := (*g)(rtn)
	if readgstatus(gp) == waiting {
		goready(gp, 1)
	}
}
