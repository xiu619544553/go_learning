package main

import (
	"fmt"
)

func aboutForTest() {

	fmt.Println("===for循环===")
	
	// Go语言开发中只有一种循环for, 比较简洁。
	// 回到Go语言中的for循环，Go语言中的for有个与其它的不一样的是for后面可以不带初始化或条件语句，即使有语句也不需要使用括号，直接空格写语句即可
	// GO语言示例如下：
	// 第一种写法：就是直接for{}，所有的条件判断逻辑都写在{}中，自己根据条件控制break逻辑。
	// 第二种写法：for后面只跟一个条件判断语句，其它的逻辑写在{}中
	// 第三种写法：和其它的语言就比较像，初始化，条件, 逻辑{}
	
	// forTest1()
	// forTest2()
	// forTest3()
}

const height = 10

// 第一种写法：就是直接for{}，所有的条件判断逻辑都写在{}中，自己根据条件控制break逻辑。
func forTest1() {
	var j = 1
	for {
		if j >= height {
			// 退出 for 循环
			break
		}
		j ++
		fmt.Println("j = ", j)
	}
}

// 第二种写法：for后面只跟一个条件判断语句，其它的逻辑写在{}中
func forTest2() {
	j := 1
	for j < height {
		j ++
		fmt.Println("j = ", j)
	}
}

// 第三种写法：和其它的语言就比较像，初始化，条件, 逻辑{}
func forTest3() {
	for i := 0; i < height; i++ {
		fmt.Println("i = ", i)
	}
}