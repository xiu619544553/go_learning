package main

import (
	"fmt"
	"time"
)

func aboutChanTimeout() {
	fmt.Println("===channel 超时机制===")

	// 与 switch 语句相比，select 有比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个 IO 操作，大致的结构如下：
	timeoutTest()
}


func timeoutTest() {

	ch := make(chan int)
	quit := make(chan bool)

	// 新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()
		
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("程序结束")
}