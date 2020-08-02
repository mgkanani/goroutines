package goroutines

import "unsafe"

func goRoutineID() int64

func goRoutine() *g

// CurRoutine returns current go routine pointer.
func CurRoutine() unsafe.Pointer {
	return unsafe.Pointer(goRoutine())
}

/*
// TODO
func CurRoutineID() int64 {
	return goRoutineID()
}
*/
