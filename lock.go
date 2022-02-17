package main

import (
	"fmt"
	"sync"
)

func aboutLock() {
	fmt.Println("===锁===")

	// Go语言包中的 sync 包提供了两种锁类型：sync.Mutex 和 sync.RWMutex。

	// sync.Mutex
	// Mutex 是最简单的一种锁类型，同时也比较暴力，当一个 goroutine 获得了 Mutex 后，其他 goroutine 就只能乖乖等到这个 goroutine 释放该 Mutex。

	// sync.RWMutex
	// RWMutex 相对友好些，是经典的单写多读模型。
	// 在读锁占用的情况下，会阻止写，但不阻止读，也就是多个 goroutine 可同时获取读锁（调用 RLock() 方法；
	// 而写锁（调用 Lock() 方法）会阻止任何其他 goroutine（无论读和写）进来，整个锁相当于由该 goroutine 独占。从 RWMutex 的实现看，RWMutex 类型其实组合了 Mutex：
	


	// Mutex 读取操作
	lockTest1()

	// RWMutex 互斥锁
}

// 读取操作
func lockTest1() {
	// 可以进行并发安全的设置
	SetCount(1)

	// 可以进行并发安全的获取
	fmt.Println(GetCount())
}

// 定义变量
var (
	// 逻辑中使用的某个变量
	count int

	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
)

// 读
func GetCount() int {
	// 锁定
	countGuard.Lock()

	// 在函数退出时接触锁定
	countGuard.Unlock()

	return count
}

// 写
func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}


// RWMutex 互斥锁
func lockTest2() {
	
}

var (
	num int
	// 与变量对应的使用互斥锁
	numGuard sync.RWMutex
)

func GetNum() int {
	// 锁定
	numGuard.RLock()

	// 在函数退出时接触锁定
	numGuard.RUnlock()

	return count
}