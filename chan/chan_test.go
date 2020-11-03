package chann

import (
	"testing"
)

func TestClose(t *testing.T) {
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 0)
	go Worker(ch1,ch2)
	close(ch1)
	
	// time.Sleep(3*time.Second)
	<-ch2
}
