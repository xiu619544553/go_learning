package main

import (
	"fmt"
	"runtime"
)

func aboutConcurrent() {
	fmt.Println("===多核并行化===")

	// 多核并行
	concurrentTest1()

	
}

// 多核并行
func concurrentTest1() {
	for i := 0; i < 5; i++ {
		go AsyncFunc(i)
	}

	
	// 2秒睡眠。保证 main.goroutine 不会退出。
	// time.Sleep(2 * time.Second)


	// 获取当前设备的cpu核心数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu核心数：", cpuNum)


	// 虽然Go语言还不能很好的利用多核心的优势，我们可以先通过设置环境变量 GOMAXPROCS 的值来控制使用多少个 CPU 核心。
	// 具体操作方法是通过直接设置环境变量 GOMAXPROCS 的值，或者在代码中启动 goroutine 之前先调用以下这个语句以设置使用 16 个 CPU 核心：
	// 设置需要用到的cpu数量
	// runtime.GOMAXPROCS(cpuNum)
}

func AsyncFunc(index int) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += 1
	}
	fmt.Printf("线程%d，sum为:%d\n", index, sum)
}


// 模拟并行的计算任务

type Vector []float64

// 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for i := 0; i < n; i++ {
		// v[i] += u.Op(v[i])
	}
}
