package channel

import (
	"fmt"
	"os"
	"time"
)

var ch = make(chan int,10)
var ch2 = make(chan string,10)

func ChanHandler() {
	time.Sleep(time.Second)
	for {
		select {
		case intVal := <-ch:
			fmt.Println("Got integer", intVal)
		case strVal := <-ch2:
			fmt.Println("Got string", strVal)
		case <-stopper:
			close(ch)
			close(ch2)
			close(stopper)
			os.Exit(0)
		default:
			fmt.Println("Nothing in channels")
			time.Sleep(time.Second)
		}
	}
}

func F3() {
	go ChanHandler()
	for i:=0 ;i<10; i++ {
		ch<-i*5
		ch2<-"String " + fmt.Sprint(i*100)
	}
	time.Sleep(time.Second*2)
	stopper<-true
}
