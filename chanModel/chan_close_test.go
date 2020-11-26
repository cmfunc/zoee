package chanModel

import (
	"fmt"
	"testing"
	"time"
)

// TestChanCloseDefault 将一个未消费完的channel关闭
func TestChanCloseDefault(t *testing.T) {

	ChanData = make(chan int, 7)

	quit:=make(chan struct{})
	go Subscribe(quit)

	go func(){
		for{
			time.Sleep(time.Millisecond*500)
			fmt.Println("ChanData 长度",len(ChanData))
		}
	}()

	for i := 0; i < 7; i++ {
		Publish(i)
	}
		
	fmt.Println("执行关闭ChanData前")
	close(ChanData)
	fmt.Println("执行关闭ChanData后")


	<-quit
}
