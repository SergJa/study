package channel

import (
	"errors"
	"fmt"
	"time"
)

// Concurrent integer
type CInt struct {
	data int
	// канал передачи на чтение
	read chan int
	// канал запроса на чтение
	getRead chan bool
	// канал записи
	write chan int

	writeResult chan error
	// сигнальный канал для остановки обработчика
	closer chan bool
}

func NewCInt() *CInt {
	c := CInt{}
	// инициализация каналов
	c.read = make (chan int)
	c.getRead = make (chan bool)
	c.write = make(chan int)
	c.writeResult = make (chan error)
	c.closer = make(chan bool)
	go c.chanHandler()
	return &c
}

func (c* CInt) Close() {
	// сигнал о завершении хендлеру
	c.closer<-true
	// закрытие каналов
	close(c.read)
	close(c.write)
	close(c.getRead)
	close(c.writeResult)
	close(c.closer)
}

func (c* CInt) chanHandler() {
	for {
		select {
		case <-c.getRead:
			// логика обработки чтения
			c.read <- c.data
		case w := <-c.write:
			// логика обработки записи
			if w > 0 {
				c.data = w
				c.writeResult <- nil
			} else {
				c.writeResult <- errors.New("Must be greater zero")
			}
		case <-c.closer:
			// закрытие обработчика
			break
		}
	}
}

func (c* CInt) Write(newValue int) error {
	c.write <- newValue
	return  <-c.writeResult
}

func (c* CInt) Read() int {
	c.getRead<-true
	return <-c.read
}

func main() {
	var cInt = NewCInt()
	defer cInt.Close()
	for i:=0; i<1000; i++ {
		go func() {
			cInt.Write(i)
			fmt.Println(cInt.Read(), "=", cInt.data)
		}()
	}
	time.Sleep(time.Second)
}

