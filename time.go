package goroutines

//requires dynamic linking to work
import _ "unsafe"

// Now() can be used for metrics recording, faster than time.Now().
// Since it uses local processor ts, it may not be accurate.
// However, metric related use cases doesn't require ns granularity.
//go:linkname Now runtime.nanotime
func Now() int64
