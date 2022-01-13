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
	fmt.Println(f, g)

	// 匿名变量
	// 匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
	h, _ := getData()
	fmt.Println(h)


	// 常量
	// 常量在定义的时候必须赋值，定义程序运行期间不会改变的那些值
	const pi = 3.14

	// iota

	fmt.Println("==========")
	// 多行字符串
	// 定义多行字符串时，必须使用反引号字符
	s1 := `第一行
	第二行
	第三行`
	fmt.Println(s1)

	// 计算字符串长度
	s1Len := len(s1)
	fmt.Println(s1Len)



	// rune：Unicode Code Point, int32
	c1 := 'x'
	fmt.Println(c1)
}

func getData() (int, int) {
	return 100, 200
}