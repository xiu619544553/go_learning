package main

import (
	"fmt"
)

func aboutBufferedChan() {
	fmt.Println("===buffered channel===")

	// 1、创建带缓冲通道
	// 通道实例 := make(chan 通道类型, 缓冲大小)
	// * 缓冲大小：决定通道可以保存的元素数量
}

// 1、创建带缓冲通道
func bufferedTest1() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(len(ch))
}
