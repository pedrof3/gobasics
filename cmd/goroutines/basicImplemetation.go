package goroutines

import (
	"fmt"
	"time"
)

func Printer() {
	chl := make(chan bool)
	chc := make(chan bool)

	go linux(chl)
	go cybersecurity(chc)

	<-chl
	<-chc
}

func linux(ch chan<- bool) {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d Linux\n", i)
		time.Sleep(time.Millisecond * 10)
	}
	ch <- true
	close(ch)
}

func cybersecurity(ch chan<- bool) {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d Cybersecurity\n", i)
		time.Sleep(time.Millisecond * 15)
	}
	ch <- true
	close(ch)
}
