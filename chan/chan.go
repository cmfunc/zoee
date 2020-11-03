package chann

import "fmt"


// Worker 关闭
func Worker(ch1,ch2 chan int){
	select {
	case v,ok:=<-ch1:
		if !ok{
			fmt.Println("接收到关闭的chan",v)
		}
		fmt.Println(v)
	default:
		fmt.Println("默认")
	}
	ch2<-5
}