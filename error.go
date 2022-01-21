package main

import (
	"fmt"
)

// 异常处理
func aboutError() {

	/*
	panic：
	    1、内置函数
		2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
		3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
		4、直到goroutine整个退出，并报告错误
	*/

	/*
	recover：
	    1、内置函数
		2、用来控制一个 goroutine 的 panicking 行为，捕获 panic，从而影响应用的行为
		3、一般的调用建议
			a). 在 defer 函数中，通过 recever 来终止一个 goroutine 的 panicking 过程，从而恢复正常代码的执行
			b). 可以获取通过 panic 传递的 error
	*/

	/*
	注意：
		1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
    	2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
    	3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。
	*/

	// errorTest1()

	// 向已关闭的通道发送数据，引发 panic
	errorTest2()
}


func errorTest1() {
	defer func() {
		fmt.Println("o NO")
		if err := recover(); err != nil {
			fmt.Println(err.(string)) // 将 interface{} 转型为具体类型
		}
	}()
	
	panic("panic error!")	
}


// 向已关闭的通道发送数据，引发 panic
func errorTest2() {
	defer func ()  {
		if err := recover(); err != nil {
			// 输出结果：
			// send on closed channel
			fmt.Println(err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}



