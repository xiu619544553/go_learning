package main

import (
	"fmt"
)

func aboutDeferClosure() {
	fmt.Println("===闭包，Defer，参数传递，变量作用域===")

	f := addUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16
}

// 累加器
func addUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(i int) int {
		n += i
		str = fmt.Sprintf("%s_$", str)
		fmt.Println("str=", str) // str= hello_$、str= hello_$_$、str= hello_$_$_$
		return n
	}
}

/*
对上面代码的说明和总结：
	1、addUpper 是一个函数，返回值的数据类型是 func(int) int
	2、闭包的说明：

	var n int = 10
	return func(i int) int {
		n += i
		return n
	}

	返回的是一个匿名函数，但是这个匿名函数引用到函数外的变量 n，因此这个匿名函数就和 n 形成一个整体，构成闭包
	
	3、可以这么理解：闭包是类，函数是操作，n 是属性字段。函数和它使用到的 n 组成了闭包。
	4、当反复的调用 f 函数时，因为 n 是初始化一次，因为每调用一次就进行累加
	5、搞清楚闭包的关键，就是要分析出返回的函数它使用（引用）到哪些变量，因为函数和它引用的变量共同构成闭包。
*/