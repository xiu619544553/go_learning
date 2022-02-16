package main

import (
	"math/rand"
	"fmt"
	"sync"
)

func aboutUnBufferedChannel() {
	fmt.Println("===Go语言无缓冲的通道===")

	// 模拟网球比赛
	bcTest1()
}


// 模拟网球比赛
func bcTest1() {

	// 创建无缓冲通道
	count := make(chan int)

	// 计数 +2，表示要等待两个 goroutine
	cwg.Add(2)

	// 启动两个选手
	go player("Alex", count)
	go player("Jack", count)

	// 发球
	count <- 1

	// 等待游戏结束
	cwg.Wait()
}

var cwg sync.WaitGroup

// player 模拟一个选手在打网球
func player(name string, count chan int) {

	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer cwg.Done()
	
	for {

		// 等待球被击打过来
		ball, ok := <-count
		
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道，表示我们输了
			close(count)
			return
		}

		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		count <- ball
	}
}