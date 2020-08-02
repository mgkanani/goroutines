package goroutines

import (
	"sync/atomic"
	"unsafe"
)

// Producer can be either Single or Multi
type Producer interface {
	Produce(item interface{})
	Close()
}

// Consumer can be either Single or Multi
type Consumer interface {
	Consume() interface{}
	Close()
}

// SPSC follows Single Producer Single Consumer patterned
type SPSC interface {
	Produce(item interface{})
	Consumer
}

type spsc struct {
	cap        int //const
	ringBuffer []interface{}
	head       int
	tail       int
	len        int64
	blockedRtn unsafe.Pointer // unsafe.Pointer(*g)
	isClosed   bool
	mutex      *mutex
}

// NewSPSC create a SingleProducerSingleConsumer object and returns
func NewSPSC(cap int) SPSC {
	pc := &spsc{
		ringBuffer: make([]interface{}, cap),
		head:       0,
		tail:       0,
		len:        0,
		cap:        cap,
		mutex:      &mutex{},
	}
	atomic.StorePointer(&pc.blockedRtn, nil) //.Store(nil)
	return pc
}

// Close notifies Producer/Consumer that Peer won't send data anymore.
func (ring *spsc) Close() {
	lock(ring.mutex)
	if ring.isClosed {
		panic("closing Closed List!")
	}
	ring.isClosed = true
	if ring.blockedRtn != nil {
		Awake(ring.blockedRtn)
		ring.blockedRtn = nil
	}
	unlock(ring.mutex)
}

func (ring *spsc) unblockConsumer(rtn *g, lck unsafe.Pointer) bool {
	if ring.blockedRtn != nil {
		Awake(ring.blockedRtn)
	}
	if ring.isClosed {
		panic("Produce can not be performed on a closed List!")
	}
	ring.blockedRtn = unsafe.Pointer(rtn)
	unlock((*mutex)(lck))
	return true
}

func (ring *spsc) unblockProducer(rtn *g, lck unsafe.Pointer) bool {
	if ring.blockedRtn != nil {
		Awake(ring.blockedRtn)
	}
	if ring.isClosed {
		unlock((*mutex)(lck))
		// don't allow consumer to block
		return false
	}
	ring.blockedRtn = unsafe.Pointer(rtn)
	unlock((*mutex)(lck))
	return true
}

func (ring *spsc) Produce(elem interface{}) {
	if atomic.LoadInt64(&ring.len) == int64(ring.cap) {
		//block if ring full
		Park(true, ring.unblockConsumer, ring.mutex)
	}

	ring.ringBuffer[ring.head] = elem
	ring.head = (ring.head + 1) % ring.cap
	atomic.AddInt64(&ring.len, 1)
}

func (ring *spsc) Consume() (elem interface{}) {

	if atomic.LoadInt64(&ring.len) == int64(0) {
		//block if ring is empty
		Park(false, ring.unblockProducer, ring.mutex)
	}

	ret := ring.ringBuffer[ring.tail]
	ring.tail = (ring.tail + 1) % ring.cap
	atomic.AddInt64(&ring.len, -1)
	return ret
}
