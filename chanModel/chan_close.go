package chanModel

import (
	"fmt"
	"time"
)

var ChanData chan int 

func Publish(in int) {
	ChanData <- in
}

func Subscribe(quit chan struct{}) {

	i:= 0
	for {
		if i==10{
			quit<-struct{}{}
		}
		select {
		case i, ok := <-ChanData:
			if ok {
				fmt.Println("接收到ChanData 非零chan数据",i)
			}else{
				fmt.Println("ChanData 已关闭")
			}
		
			
		// default:
		// 	fmt.Println("执行默认default")
		}
		time.Sleep(time.Second*2)
		i++
	}


}
