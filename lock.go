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
	// lockTest1()


	// RWMutex 互斥锁
	// lockTest2()


	// 死锁、活锁、饥饿锁
	// 1、死锁
	// 死锁是指两个或两个以上的进程（或线程）在执行过程中，因争夺资源而造成的一种互相等待的现象，若无外力作用，它们都将无法推进下去。此时称系统处于死锁状态或系统产生了死锁，这些永远在互相等待的进程称为死锁进程。

	// 死锁发生的条件有如下几种：
	// ① 互斥条件
	// ② 请求和保持条件
	// ③ 不剥夺条件
	// ④ 环路等待条件

	// 死锁解决办法：
	//  	如果并发查询多个表，约定访问顺序；
	//		在同一个事务中，尽可能做到一次锁定获取所需要的资源；
	//		对于容易产生死锁的业务场景，尝试升级锁颗粒度，使用表级锁；
	//		采用分布式事务锁或者使用乐观锁。


	// 2、活锁
	// 活锁是另一种形式的活跃性问题，该问题尽管不会阻塞线程，但也不能继续执行，因为线程将不断重复同样的操作，而且总会失败。
	// 活锁通常发生在处理事务消息中，如果不能成功处理某个消息，那么消息处理机制将回滚事务，并将它重新放到队列的开头。
	// 这样，错误的事务被一直回滚重复执行，这种形式的活锁通常是由过度的错误恢复代码造成的，因为它错误地将不可修复 的错误认为是可修复的错误。


	// 3、饥饿
	// 饥饿是指一个可运行的进程尽管能继续执行，但被调度器无限期地忽视，而不能被调度执行的情况。
	// 与死锁不同的是，饥饿锁在一段时间内，优先级低的线程最终还是会执行的，比如高优先级的线程执行完之后释放了资源。
	// 更广泛地说，饥饿通常意味着有一个或多个贪婪的并发进程，它们不公平地阻止一个或多个并发进程，以尽可能有效地完成工作，或者阻止全部并发进程。
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