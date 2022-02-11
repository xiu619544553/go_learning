package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func aboutGoRoutine() {
	fmt.Println("===goroutine===")
	// 并发和并行
	//     A. 多线程程序在一个核的cpu上运行，就是并发。
	//     B. 多线程程序在多个核的cpu上运行，就是并行。


	// 进程和线程
	// 进程是资源分配的最小单位，线程是CPU调度的最小单位


	// 协程和线程
	// 协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。
	// 线程：一个线程上可以跑多个协程，协程是轻量级的线程。


	// 在程序启动时，Go 程序就会为 main() 函数创建一个默认的 goroutine。


	// goroutine 只是由官方实现的超级"线程池"。
	// 每个实例 4~5KB 的栈内存占用和由于实现机制而大幅减少的创建和销毁开销是go高并发的根本原因。


	// 并发不是并行：
	// 并发主要由切换时间片来实现"同时"运行，并行则是直接利用多核实现多线程的运行，go可以设置使用核数，以发挥多核计算机的能力。

	// goroutine 奉行通过通信来共享内存，而不是共享内存来通信。




	// 使用goroutine
	// Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。
	// 一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。


	// 启动单个goroutine
	// 启动goroutine的方式非常简单，只需要在调用的函数（普通函数和匿名函数）前面加上一个go关键字。


	// 测试1 -- 串行
	// routineTest1()

	// 测试2 -- 启动单个 goroutine
	// routineTest2()

	// 启动多个goroutine
	// routineTest3()

	// 测试4
	// routineTest4()

	// 测试5 -- 每隔一秒打印一次计数器
	// routineTest5()

	// 处理线程间数据共享的
	// routineTest6()

	// Go语言提供的是另一种通信模型，即以消息机制而非共享内存作为通信方式。
	// Go语言提供的消息通信机制被称为 channel

	// 资源竞争
	// routineTest7()

	// runtime.Gosched 理解
	// routineTest8()

	// 锁住共享资源
	// routineTest9()

	// 互斥锁
	routineTest10()
}


// 测试1
func routineTest1() {
	// 这个示例中hello函数和下面的语句是串行的
	rHello()
	fmt.Println("main goroutine done!")
}

func rHello() {
	fmt.Println("Hello Goroutine!")
}


// 启动单个goroutine
func routineTest2() {
	// 首先为什么会先打印main goroutine done!是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的。
	go rHello2()
	fmt.Println("main goroutine done!2")
	time.Sleep(time.Second) // 等待几秒，否则 rHello2不会执行。这是因为：main函数执行完毕后，主协程退出了，所以 rHello2 不会执行了。
}

func rHello2() {
	fmt.Println("Hello Goroutine2!")
}


// 启动多个goroutine

var wg sync.WaitGroup

func routineTest3() {
	fmt.Println("===启动多个goroutine===")

	// 多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。

	for i := 0; i < 10; i++ {
		wg.Add(1)  // 启动一个 goroutine 就登记 +1
		go rTest3(i)
	}

	wg.Wait() // 等待所有登记的 goroutine 都结束
}

func rTest3(i int) {
	defer wg.Done() // goroutine 结束就登记 -1
	fmt.Println("Hello Goroutine!", i)
}


// 测试4
func routineTest4() {
	// 合起来写
    go func() {
        i := 0
        for {
            i++
            fmt.Printf("new goroutine: i = %d\n", i)
            time.Sleep(time.Second)
        }
    }()
    i := 0
    for {
        i++
        fmt.Printf("main goroutine: i = %d\n", i)
        time.Sleep(time.Second)
        if i == 2 {
            break
        }
    }
}

// 测试5
func routineTest5() {
	// 为一个普通函数创建 goroutine 的写法
	// go 函数名(参数列表)
	// 注：使用 go 关键字创建 goroutine 时，被调用函数的返回值会被忽略。
	// 如果需要在 goroutine 中返回数据，请使用后面介绍的通道（channel）特性，通过通道把数据从 goroutine 中作为返回值传出。
	// 提示：所有的 goroutine 在 main() 函数结束时会一同结束

	// 并发执行程序
	go running()

	// 为匿名函数创建 goroutine 函数
	// go func() {
	// 	var times int
	// 	for {
	// 		times++
	// 		fmt.Println("tick", times)

	// 		// 延时1秒
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	// 接受命令行输入，不做任何事情
	var input string
	fmt.Scanln(&input)
}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)

		// 延时1秒
		time.Sleep(time.Second)
	}
}


// 处理线程间数据共享的
func routineTest6() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

var counter int = 0
func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("counter = ", counter)
	lock.Unlock()
}

// 资源竞争
func routineTest7() {
	// 有并发，就有资源竞争，如果两个或者多个 goroutine 在没有相互同步的情况下，访问某个共享的资源，比如同时对该资源进行读写时，就会处于相互竞争的状态，这就是并发中的资源竞争。

}


// runtiime.Gosched 的理解
// runtime.Gosched()用于让出CPU时间片。这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A歇着了，B继续跑。
// runtime.Gosched() 是让当前 goroutine 暂停的意思，退回执行队列，让其他等待的 goroutine 运行
func routineTest8() {
	go rSay("world")
	rSay("hello")
}

func rSay(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Printf("i = %d，s = %s\n", i, s)
	}
}


// 锁住共享资源
// Go语言提供了传统的同步 goroutine 的机制，就是对共享资源加锁。atomic 和 sync 包里的一些函数就可以对共享的资源进行加锁操作。
func routineTest9() {
	swg.Add(2)
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 结束
	swg.Wait()
	fmt.Println(number)
}

var (
	number int64
	swg sync.WaitGroup
)
func incCounter(id int) {
	defer swg.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt64(&number, 1) // 安全的对 number 加1
		runtime.Gosched()
	}
}


// 互斥锁
// 另一种同步访问共享资源的方式是使用互斥锁，互斥锁这个名字来自互斥的概念   http://c.biancheng.net/view/4358.html
func routineTest10() {
	sg.Add(2)

	go incLength(1)
	go incLength(2)

	sg.Wait()
	fmt.Println(length)
}

var (
	length int64
	sg sync.WaitGroup
	mutex sync.Mutex
)
func incLength(id int) {
	defer sg.Done()
	
	for i := 0; i < 2; i++ {

		// 同一时刻，只允许一个 goroutine 进入这个临界区
		mutex.Lock()
		{
			value := length
			runtime.Gosched()
			value++
			length = value
		}

		// 释放锁，允许其他正在等待的 goroutine 进入临界区
		mutex.Unlock() 
	}
}

func routineTest11() {
	
}