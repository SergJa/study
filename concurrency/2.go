package concurrency

import (
	"fmt"
	"sync"
)

var concurrentInt int

var mx sync.Mutex

func P2(i int) {
	mx.Lock()
	defer mx.Unlock()
	fmt.Println("Previous", concurrentInt, "New", i)
	concurrentInt = i
	wg.Done()
}

//var wg sync.WaitGroup

func F2() {
	routineCount := 1000
	wg.Add(routineCount)

	for i:=0; i<routineCount; i++ {
		go P2(i)
	}
	wg.Wait()
}
