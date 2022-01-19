package main

import (
	"fmt"
)

// 闭包
func aboutClosure() {
	
	// 在闭包内部修改引用的变量
	aboutClosure1()

	// 闭包的记忆效应
	aboutClosure2()

	// 闭包生成器使用场景
	closureUsageScene()

	// 闭包复制的是原对象指针
	aboutClosureObjPrt()
}

// 在闭包内部修改引用的变量
func aboutClosure1() {

	str := "hello world"
	fmt.Printf("修改前：str = %s\n", str)
	
	// 匿名函数
	foo := func() {
		str = "hello go"
	}

	// 调用匿名函数
	foo()

	fmt.Printf("修改后：str = %s\n", str)
}

// 闭包的记忆效应：被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。
func aboutClosure2() {
	
	fmt.Println("======= 闭包1 =======")
	// 创建一个累加器，初始值为 1
	accumulator1 := Accumulate(1)

	// 调用闭包
	fmt.Println(accumulator1())
	fmt.Println(accumulator1())

	// 打印闭包地址
	fmt.Printf("闭包1地址：%p\n", &accumulator1)
	

	fmt.Println("======= 闭包2 =======")
	// 新创建一个累加器，初始值为 10
	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Println(accumulator2())
	fmt.Printf("闭包2地址：%p\n", &accumulator2)
}

// 该函数每调用一次，会对value累加一次。函数返回值是闭包
func Accumulate(value int) func() int {
	return func() int {

		// 累加
		value ++

		// 返回累加的值
		return value
	}
}



/************************* 闭包的记忆效应使用场景 **************************/

func closureUsageScene() {
	// 创建一个玩家生成器
	generator := playerGen("Tom")

	name, hp := generator()

	fmt.Printf("name=%s，hp=%d\n", name, hp)
}

// 闭包的记忆效应使用场景：实现类似于设计模式中工厂模式的生成器
func playerGen(name string) func() (string, int) {

	// 人物的血量
	hp := 150

	// 返回闭包函数
	return func() (string, int) {

		// 将变量引用到闭包中
		return name, hp
	}
}

/************************* 闭包复制的是原对象指针 **************************/

func aboutClosureObjPrt() {

	f := closure3(1)

	f()

	/*
	输出结果：
		x的值：1，地址：0xc0000160b8
		x的值：1，地址：0xc0000160b8
	闭包复制的是原对象的指针

	在汇编层 closure3 实际返回的是 FuncVal 对象，其中包含了匿名函数地址、闭包对象指针。当调 匿名函数时，只需以某个寄存器传递该对象即可。
	*/
}

func closure3(x int) func() {
	fmt.Printf("x的值：%d，地址：%p\n", x, &x)

	return func() {
		fmt.Printf("x的值：%d，地址：%p\n", x, &x)
	}
}

/************************* 闭包复制的是原对象指针 **************************/