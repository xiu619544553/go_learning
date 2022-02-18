package main

import (
	"fmt"
	"net/http"
	"sync"
)

func aboutWaitGroup() {
	fmt.Println("===sync.WaitGroup===")

	// 并发线程同步的方式
	// 1、通道 channel
	// 2、互斥锁 sync.Mutex、sync.RWMutex
	// 3、等待组 sync.WaitGroup


	// var group sync.WaitGroup
	// group.Add(1)
	// group.Done()
	// group.Wait()


	// sync.WaitGroup 初步使用方式
	// groupTest1()


	// 死锁解决办法：
	// 	如果并发查询多个表，约定访问顺序；
	// 	在同一个事务中，尽可能做到一次锁定获取所需要的资源；
	// 	对于容易产生死锁的业务场景，尝试升级锁颗粒度，使用表级锁；
	// 	采用分布式事务锁或者使用乐观锁。

	
}

func groupTest1() {
	
	// 声明一个等待组
	var group sync.WaitGroup

	var urls = []string{
		"https://www.github.com",
		"https://www.baidu.com",
		"https://www.golangtc.com",
	}

	for _, url := range urls {
		
		// 每一个任务开始时，将等待组 +1
		group.Add(1)

		// 开启一个协程
		go func(url string) {
			
			// 使用defer，表示函数完成时将等待组 —1
			defer group.Done()

			// 使用http访问提供的地址
			_, err := http.Get(url)

			// 打印结果
			fmt.Println(url, err)

			// 通过参数传递url地址
		}(url)
	}

	// 等待所有的任务完成
	group.Wait()

	fmt.Println("=== Over ===")
}

