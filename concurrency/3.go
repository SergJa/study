package concurrency

import (
	"sync/atomic"
)

var concurrentInt32 int32

func F3() {
	atomic.AddInt32(&concurrentInt32, 5)
}
