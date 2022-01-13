package main // 声明 main 包

import (
	"fmt" // 导入 fmt 包，打印字符串是需要用到
	"strings"
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


	fmt.Println("==========常量==========")
	// 常量
	// 常量在定义的时候必须赋值，定义程序运行期间不会改变的那些值
	const pi = 3.14

	// iota

	fmt.Println("==========字符串==========")
	// 多行字符串
	// 定义多行字符串时，必须使用反引号字符
	s2 := `第一行
	第二行
	第三行`
	fmt.Println(s2)

	// 单行字符串
	s1 := "第一行第二行第三行"

	// 计算字符串长度
	s1Len := len(s1)
	fmt.Println(s1Len)

	if strings.Contains(s1, "第") {
		fmt.Println("s1包含有字符\"第\"")
	}
	
	// 拼接字符串
	// s1 = fmt.Sprintf("\n第四行") 

	// 分割
	var sr []string = strings.Split(s1, "第") 
	fmt.Println("sr = ", sr)

	// 前缀/后缀判断
	if strings.HasPrefix(s1, "第") {
		fmt.Println("s1的前缀是 第")
	}

	if strings.HasSuffix(s1, "行") {
		fmt.Println("s1的后缀是 行")
	}

	// 子串出现的位置
	subIndex := strings.Index(s1, "第")
	fmt.Println("子串出现的位置：", subIndex)

	lastIndex := strings.LastIndex(s1, "第")
	fmt.Println("子串出现的位置：", lastIndex)

	// join 操作
	// strings.Join(a[]string, sep string)	
	// joinS := strings.Join({"hello"}, s1)

	// rune：Unicode Code Point, int32
	c1 := 'x'
	fmt.Println(c1)


	fmt.Println("==========rune==========")
	// rune
	traversalString()
}

func getData() (int, int) {
	return 100, 200
}

func traversalString() {
	s := "Ha你好"
	
	count := len(s)
	fmt.Println("count =", count)
	
	for i := 0; i < len(s); i++ {
		// uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
		fmt.Printf("%d：%v(%c)", i, s[i], s[i])
		fmt.Println()
	}

	for index, r := range s {
		// rune
		fmt.Printf("%v，%v(%c)\n", index, r, r)
	}
	
}

// for range
func aboutForRange() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for i, v := range slice {
		myMap[i] = &v
	}

	fmt.Println("===new map===")
	for , v := range v {
		
	}
}