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

	fmt.Println("==========for range==========")
	aboutForRange()

	fmt.Println("==========Array==========")
	aboutArray()
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
		// ❌这是错误的
		// myMap[i] = &v

		// 创建一个新的变量，如此方能保证value对应的地址不同
		num := v
		myMap[i] = &num
	}

	fmt.Println("===new map===")
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
		// 输出结果如下，不是预期结果
		// map[1]=3
		// map[2]=3
		// map[3]=3
		// map[0]=3

		// 是什么原因导致映射所有值都相同
		// 原因分析：因为for range创建了每个元素的副本，而不是直接返回每个元素的引用，
		// 如果使用该值变量的地址作为指向每个元素的指针，就会导致错误，在迭代时，返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以值的地址总是相同的，导致结果不如预期。
	}
}


// 数组 Array，Golang Array和以往认知的数组有很大不同。数组：是同一种数据类型的固定长度的序列。
func aboutArray() {

	// 长度是数组类型的一部分  var arr1 [5]int 与 var arr2 [10]int 是不同的类型
	// 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。


	// 数组初始化
	var arr = [...]int{1, 2, 3} // 通过初始化值确定数组长度
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%v]=%v\n", i, arr[i])
	}

	// var arr2 [5]int = [5]int{1, 2} 				// 未初始化元素值为 0
	// var arr3 = [2]int{1, 2}
	// var strs = [5]string{3: "hello", 4: "world"} // 使用索引号初始化元素
	// var arr4 = [5]int{2: 100, 4: 200}			// 使用索引号初始化元素
	var arr5 = []


	var nums = [5]int{1, 2, 3, 4, 5}

	// 遍历数组的两种方式
	// 方式一：for 循环
	for i := 0; i < len(nums); i++ {
		fmt.Printf("num[%v]=%v\n", i, nums[i])
	}

	// 修改数组的值
	nums[0] = 100

	// 方式二：for range
	for index, value := range nums {
		fmt.Printf("nums[%v]=%v\n", index, value)
	}

	// 指针数组 [n]*T，数组指针 *[n]T。
	// 数组指针：解释：首先是一个指针，一个数组的地址
	// 指针数组：解释：首先是一个数组，存储的数据类型是指针
	// 参考：https://blog.csdn.net/weixin_40123451/article/details/122092190
}