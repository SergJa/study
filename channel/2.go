package channel

import (
	"fmt"
)

var c chan int

var stopper chan bool

func Init() {
	c = make(chan int)
	stopper = make(chan bool)
	//defer close(c)
}

func Sender() {
	c<-5
}

func Receiver() {
	fmt.Println(<-c)
	stopper <- true
}

func WaitShutdown() {
	<-stopper
	// функционал закрытия
	close(c)
	close(stopper)
}

func F2() {
	Init()
	go Receiver()
	Sender()

	WaitShutdown()
}
