package runtime

//requires dynamic linking to work
import _ "unsafe"

// Now() can be used for metrics recording, faster than time.Now().
//go:linkname Now runtime.nanotime
func Now() int64

