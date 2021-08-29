package deferTest

/*
go 的return语句不具备原子性: 可以被拆分为三条执行语句

1. return 语句后紧跟的表达式执行
2. defer 一系列函数执行
3. 执行return返回

defer后跟匿名函数，直接传参与函数体内直接引用（闭包）的区别；
直接传参的参数，在程序运行到defer函数式就已确定，并且不会改变；
直接引用defer体外的变量执行在函数return结束前，值才能确定；

*/

func Initial() (x string) {
	x = "A1"
	defer func() {
		x = "A2"
	}()

	defer func() {
		x = "A3"
	}()

	defer func() {
		x = "A4"
	}()

	defer func (m string)  {
		x = x + m
	}(x)

	x = "B"

	return returnX(x)
}

func returnX(x string) string {
	return x+"return"
}
