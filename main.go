package main // 声明 main 包

import (
	"fmt" // 导入 fmt 包，打印字符串是需要用到
)

func main () {  // 声明 main 主函数
	fmt.Println("hello world!")

	// fmt.Println("==========基本数据类型==========")
	// baseDataType()

	// fmt.Println("==========数组 Array==========")
	// aboutArray()

	// fmt.Println("==========切片 Slice==========")
	// aboutSlice()

	// fmt.Println("==========指针==========")
	// aboutPointer()

	// fmt.Println("==========Map==========")
	// aboutMap()

	// fmt.Println("==========struct==========")
	// aboutStruct()

	// fmt.Println("==========流程控制==========")
	// aboutFlow()

	// fmt.Println("==========函数==========")
	// aboutFunction()

	// fmt.Println("==========闭包==========")
	// aboutClosure()

	// fmt.Println("========== defer 延迟调用 ==========")
	// aboutDefer()

	// fmt.Println("========== 闭包，Defer，参数传递，变量作用域 ==========")
	// aboutDeferClosure()

	// fmt.Println("========== 异常处理 ==========")
	// aboutError()

	// fmt.Println("========== 方法 ==========")
	// aboutMethod()

	// fmt.Println("========== 面向对象--接口 ==========")
	// aboutInterface()

	// fmt.Println("========== http编程 ==========")
	// aboutHttp()

	// fmt.Println("========== goroutine ==========")
	// aboutGoRoutine()	

	// fmt.Println("========== channel ==========")
	// aboutChannel()

	// fmt.Println("========== 无缓冲的通道（unbuffered channel） ==========")
	// aboutUnBufferedChannel()

	// fmt.Println("=== 字符串处理 ===")
	// aboutStrings()

	// fmt.Println("========== 带缓冲的通道（buffered channel） ==========")
	// aboutBufferedChan()

	// fmt.Println("========== channel 超时机制 ==========")
	// aboutChanTimeout()

	// fmt.Println("========== 多核并行化 ==========")
	// aboutConcurrent()

	// fmt.Println("========== 互斥锁和读写互斥锁 ==========")
	// aboutLock()

	// fmt.Println("========== 等待组 ==========")
	// aboutWaitGroup()

	// fmt.Println("========== 反射 reflect ==========")
	// aboutReflect()

	fmt.Println("========== Go语言结构体标签（Struct Tag）	==========")
	aboutStructTag()
}