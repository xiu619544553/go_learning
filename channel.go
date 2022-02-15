package main

import (
	"fmt"
	"time"
)

func aboutChannel() {
	fmt.Println("==channel")

	// 参考
	// http://c.biancheng.net/view/97.html

	// 0、概述
	// 如果说 goroutine 是 Go语言程序的并发体的话，那么 channels 就是它们之间的通信机制。
	// 一个 channels 是一个通信机制，它可以让一个 goroutine 通过它给另一个 goroutine 发送值信息。
	// 每个 channel 都有一个特殊的类型，也就是 channels 可发送数据的类型。一个可以发送 int 类型数据的 channel 一般写为 chan int。


	// 1、通道的特性
	// Go语言中的通道（channel）是一种特殊的类型。在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。goroutine 间通过通道就可以通信。
	// 通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。


	// 2、声明通道类型
	// 通道本身需要一个类型进行修饰，就像切片类型需要标识元素类型。通道的元素类型就是在其内部传输的数据类型，声明如下：
	// var 通道变量 chan 通道类型
	// 通道类型：通道内的数据类型。
	// 通道变量：保存通道的变量。
	// chan 类型的空值是 nil，声明后需要配合 make 后才能使用。


	// 3、创建通道
	// 通道是引用类型，需要使用 make 进行创建，格式如下：
	// 通道实例 := make(chan 数据类型)
	// 数据类型：通道内传输的元素类型。
	// 通道实例：通过make创建的通道句柄。
	
	// chanTest1()


	// 4、使用通道发送数据
	// 通道发送数据的格式：通道发送使用特殊操作符：<-，将数据通过通道发送的格式是：
	// 通道变量 <- 值
	// 通道变量：通过make创建好的通道实例。
	// 值：可以是变量、常量、表达式或者函数返回值等。值的类型必须与ch通道的元素类型一致。
	
	// chanTest2()


	// 5、使用通道接收数据
	// 参考：http://c.biancheng.net/view/97.html 使用通道接收数据部分的内容
	// chanTest3()


	// 6、循环接收
	// chanTest4()


	// 7、单向通道的声明格式
	// var 通道实例 chan<- 元素类型   // 只能写入数据的通道
	// var 通道实例 <-chan 元素类型   // 只能读取数据的通道
	// chanTest5()
	

	// 8、time包中的单向通道
	// chanTest6()


	// 9、关闭channel
	// 关闭channel，直接使用 Go语言内置的 close() 函数即可
	// close(ch)


	// 10、如何判断一个 channel 是否已经被关闭？
	// 使用多重返回值的方式：x,ok := <-ch
	// ok值为false，则表示 ch 已经被关闭了 
}

// 创建通道
func chanTest1() {
	// 创建整数型通道
	ch1 := make(chan int)

	// 创建空接口类型的通道，可以存放任何数据类型
	ch2 := make(chan interface{})

	type Equip struct {}
	ch3 := make(chan *Equip)

	fmt.Printf("ch1类型：%T\n", ch1)
	fmt.Printf("ch2类型：%T\n", ch2)
	fmt.Printf("ch3类型：%T\n", ch3)
}

// 4、使用通道发送数据
func chanTest2() {
	ch := make(chan interface{})

	// 将 0 放入通道中
	ch <- 0

	// 将 hello 字符串放入通道中
	ch <- "hello"
}

// 5、使用通道接收数据
func chanTest3() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")

		// 通过通道，通知 main 的 goroutine 
		ch <- 0

		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	// 等待匿名 goroutine
	// 该格式，表示：接收任意数据，忽略接收的数据
	<- ch

	fmt.Println("all done")
}

// 6、循环接收
func chanTest4() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		for i := 0; i < 3; i++ {
			// 发送数据
			ch <- i

			// 每次发送完时，等待
			time.Sleep(time.Second)
		}
	}()

	// 遍历接收通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Println("data：", data)

		if data == 2 {
			break
		}
	}
}

// 7、单向通道的声明格式
func chanTest5() {
	ch := make(chan int)

	// 声明一个只能写入数据的通道类型，并赋值为ch
	var chSendOnly chan<- int = ch

	// 声明一个只能读取数据的通道类型，并赋值为ch
	var chReceiveOnly <-chan int = ch;

	fmt.Println(chSendOnly)
	fmt.Println(chReceiveOnly)
}

// 8、time包中的单向通道
func chanTest6() {
	// 返回一个计时器 timer
	timer := time.NewTimer(time.Second)
	fmt.Println(timer)
}