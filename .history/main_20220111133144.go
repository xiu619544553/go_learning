package main // 声明 main 包

import (
	"fmt"   // 导入 fmt 包，打印字符串是需要用到
)

func main () {  // 声明 main 主函数
	fmt.Println("hello world!")

	// 变量标准命名
	// 行尾不需要分号
	var age int = 10
	fmt.Println(age)

	// 变量批量命名
	var (
		a int
		b string
		c []float32
		d func() bool
		e struct {
			x int
		}
	)

	// 0  [] <nil> <nil> {0}
	fmt.Println(a, b, c, d, d, e)

	// 简短命名
	i, j := 1, 2
	fmt.Println(i, j)


	// 多重赋值
	var f int = 1
	var g int = 2

	f, g = g, f
	fmt.Println()
}