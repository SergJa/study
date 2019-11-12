package channel

import (
	"fmt"
	"sync"
)

func F1() {
	var c chan int
	c = make(chan int)
	wg := sync.WaitGroup{}


	wg.Add(1)
	go func() {
		cValue := <-c
		fmt.Println(cValue)
		wg.Done()
	}()
	c<-5

	wg.Wait()

	close(c)
}