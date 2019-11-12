package concurrency

import (
	"fmt"
	"sync"
)

func P(i int) {
	fmt.Println("Hello from goroutine", i)
	wg.Done()
}

var wg sync.WaitGroup

func F1() {
	routineCount := 1000
	wg.Add(routineCount)

	for i:=0; i<routineCount; i++ {
		go P(i)
	}
	wg.Wait()
}
