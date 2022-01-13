package main

import (
	"fmt"
	"strings"
)

func aboutSlice() {
	/*
	需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

	1. 切片是数组的一个引用，因此切片是引用类型。但本质是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。 
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string 、 var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
	*/

	// golang slice data[:6:8] 两个冒号的理解：a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x

	// 0.数组与切片区别
	diff()

	// 1. 创建切片的各种方式
	createSlice()

	// 2.切片初始化
	initSlice()

	// 3.通过 make 创建切片
	makeSlice()

	// 4.用 append 内置函数操作切片
	aboutAppend()

	// 5.超出 slice 限制
	outofBoundSlice()

	// 6.切片拷贝
	copySlice()

	// 7.遍历slice
	traverseSlice()

	// 8.切片resize（调整大小）
	resizeSlice()

	// 9.字符串和切片
	stringAndSlice()

	// 10. golang slice data[:6:8] 两个冒号的理解
	testMuli()

	// 11.数组or切片转字符串：
	arrayOrSliceToString()

	// 12.Slice底层实现
	sliceBase()
}

// 0.数组与切片区别
func diff() {

	// 数组
	var arr1 = [...]int{1, 2, 3}

	// 切片
	var slice1 = make([]int, 10, 10)

	fmt.Printf("   arr1首地址：%p\n   arr1【0】首地址：%p\n   slice首地址：%p\n   slice【0】地址%p\n", &arr1, &arr1[0], &slice1, &slice1[0])

	/*
	输出结果：
	    arr1首地址：0xc00013a000
   		arr1【0】首地址：0xc00013a000

   		slice首地址：0xc000122018
   		slice【0】地址0xc00013c000

	问题：为什么数组的首地址和他头号元素的地址相同，而切片的首地址和他头号元素不同呢？？


	数组：数组其实就一段大小固定的空间，他的大小就是在初始化时给的长度len，根据长度len和元素类型大小（例如元素为int64也就是8字节）就可以很轻易的算出空间的完整大小，
	     在完全确定后就会直接敲定无法改变。他的头号元素也是整个空间的头，地址一致。
		 正因如此，数组具有以下特点：
			1.长度固定，无法追加元素，因为空间大小就给了这么多，所以[5]int和[10]int两个数组表示的是不同的类型自然无法兼容
			2.数组是值类型的，在传递时采用值拷贝：

	切片的内存布局：https://blog.csdn.net/weixin_44938441/article/details/110424749
	切片本质就是一个结构体，他里面包含三部分：address + len + cap，
		address： 就是他指向的内部数组或数组某个地方
		len：是当前的元素个数
		cap：可容纳元素总容量大小
	也不难理解为什么cap>=len
	正是如此，切片本质上是一个引用空间，该空间和元素空间完全是两个空间，所以切片的首地址和头号元素的首地址完全不同。

	值得一提的是：通过make创建的切片和通过数组[:]分的切片是有区别的！
		1.直接var slice_2 []int = arr_1[1:3]切分数组产生的切片其实就是直接对这个数组的引用，这时当切片中的值修改时会影响到原来的数组值：
		2.而通过make创建其实本质上也是先偷偷的创建一个内部数组，然后该切片在对该数组进行操作。

	*/

}

// 1. 创建切片的各种方式
func createSlice() {
	fmt.Println("===创建切片的各种方式===")

	// 1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("s1为空")
	} else {
		fmt.Println("s1有值")
	}

	// 2. :=
	s2 := []int{}

	// 3. make
	var s3 []int = make([]int, 0)

	fmt.Println(s1, s2, s3)

	// 4. 初始化赋值
	// slice := make([]type, len, cap)
	// capacity = 2
	// len = 0
	var s4 []int = make([]int, 0, 10)
	fmt.Printf("s4 = %v\n", s4)

	// make 第二个参数是 capacity 容积
	sCount := cap(s4)
	fmt.Println(sCount)

	s5 := []int{1, 2, 3}
	fmt.Printf("s5 = %v\n", s5)

	// 5. 从数组切片
	arr1 := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	s6 = arr1[1:4] // 前包后不包，即取数组下标为 [1,4) 作为切片 
	fmt.Printf("s6 = %v\n", s6)

	fmt.Printf("s6.len = %v，cap = %v\n", len(s6), cap(s6))
}

// 2.切片初始化
func initSlice () {
	fmt.Println("===切片初始化===")
	/*
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr[start:end] 
	var slice1 []int = arr[:end]        
	var slice2 []int = arr[start:]        
	var slice3 []int = arr[:] 
	var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素	
	*/

	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice0 := arr[1:2]
	slice1 := arr[:2]
	slice2 := arr[1:]
	slice3 := arr[:]
	slice4 := arr[:len(arr) - 1]

	fmt.Printf("slice0 = %v\n", slice0)
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("slice2 = %v\n", slice2)
	fmt.Printf("slice3 = %v\n", slice3)
	fmt.Printf("slice4 = %v\n", slice4)
}

// 3.通过 make 创建切片
func makeSlice() {
	/*
	var slice []type = make([]type, len)
    slice  := make([]type, len)
    slice  := make([]type, len, cap)

	切片的长度 len，总是 <= cap(slice)
	切片的容量 cap，总是 >= len(slice)
	*/
	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
    fmt.Println(s1, len(s1), cap(s1))

    s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
    fmt.Println(s2, len(s2), cap(s2))

    s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
    fmt.Println(s3, len(s3), cap(s3))

	s4 := []int{0, 1, 2, 3}
	p := &s4[2]
	*p += 100
	fmt.Printf("s4 = %v\n", s4) // s4 = [0 1 102 3]
}

// 4.用 append 内置函数操作切片
func aboutAppend() {
	var a = []int{1, 2, 3}
	fmt.Printf("a = %v\n", a)

	var b = []int{4, 5, 6}
	fmt.Printf("b = %v\n", b)

	c := append(a, b...)
	fmt.Printf("c = %v\n", c)

	d := append(c, 7)
	fmt.Printf("d = %v\n", d)

	aboutAppend2()
}

// append ：向 slice 尾部添加数据，返回新的 slice 对象。
func aboutAppend2() {
	s1 := make([]int, 0, 5)
    fmt.Printf("&s1 = %p\n", &s1)

    s2 := append(s1, 1)
    fmt.Printf("&s2 = %p\n", &s2)

    fmt.Println(s1, s2)
}

// 5.超出 slice 限制
func outofBoundSlice() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
    s := data[:2:3] // 从【0,2) 位置获取切片，len = 2 - 0;  cap = 3 - 0;

	fmt.Println("===append 前===")
	fmt.Printf("data = %v\n", data)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("s.len = %v，cap = %v\n", len(s), cap(s)) // s.len = 2，cap = 3
	fmt.Println(&s[0], &data[0])	// s、data 的底层数组起始指针时相同的

	fmt.Println("===append 后===")
	s = append(s, 100, 200)  // 一次append两个元素，超出了 s.cap 限制
	fmt.Printf("data = %v\n", data)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("s.len = %v，cap = %v\n", len(s), cap(s)) // s.len = 4，cap = 6。容积翻倍了。

	fmt.Println(&s[0], &data[0]) 	// 比对底层数组起始指针。不同。


	// slice中cap重新分配规律：原cap x 2
}

// 6.切片拷贝
func copySlice() {
	fmt.Println("===切片拷贝===")
	s1 := []int{1, 2, 3, 4, 5}
    fmt.Printf("slice s1 : %v\n", s1)		 // [1 2 3 4 5]

    s2 := make([]int, 10)
    fmt.Printf("slice s2 : %v\n", s2)	     // [0 0 0 0 0 0 0 0 0 0]
    copy(s2, s1)
    fmt.Printf("copied slice s1 : %v\n", s1) // [1 2 3 4 5]
    fmt.Printf("copied slice s2 : %v\n", s2) // [1 2 3 4 5 0 0 0 0 0]

    s3 := []int{1, 2, 3}
    fmt.Printf("slice s3 : %v\n", s3)

    s3 = append(s3, s2...)
    fmt.Printf("appended slice s3 : %v\n", s3)

    s3 = append(s3, 4, 5, 6)
    fmt.Printf("last slice s3 : %v\n", s3)

	fmt.Println("===切片拷贝2===")
	copySlice2()
}

// copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。
func copySlice2() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println("array data : ", data)

    s1 := data[8:]
    s2 := data[:5]
    fmt.Printf("slice s1 : %v\n", s1)
    fmt.Printf("slice s2 : %v\n", s2)
	
    copy(s2, s1)
    fmt.Printf("copied slice s1 : %v\n", s1)
    fmt.Printf("copied slice s2 : %v\n", s2)
    fmt.Println("last array data : ", data)
}

// 7.遍历slice
func traverseSlice() {
	fmt.Println("===遍历slice===")
	data := [...]int{1, 2, 3, 4, 5, 6}
	slice := data[:]
	for index, value := range slice {
		fmt.Printf("slice[%v]=%v\n", index, value)
	}
}

// 8.切片resize（调整大小）
func resizeSlice() {
	var a = []int{1, 3, 5, 7}
	fmt.Printf("slice a : %v，len(a) : %v\n", a, len(a))

	b := a[1:2]
	fmt.Printf("slice b : %v，len(b) : %v\n", b, len(b))

	c := b[0:3]
	fmt.Printf("slice c : %v，len(c) : %v\n", c, len(c))
}

// 9.字符串和切片
func stringAndSlice() {
	fmt.Println("===字符串和切片===")
	// string底层就是一个byte的数组，因此，也可以进行切片操作。
	str := "abcdefghijk"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)


	// string本身是不可变的，因此要改变string中的字符。需要如下操作： 

	// 9.1 英文字符串
	fmt.Println("===英文字符串 []byte(str)===")
	str1 := "Hello world"
	s := []byte(str1)	// 中文字符需要用 []rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '~')
	str1 = string(s)
	fmt.Printf("str1 = %v\n", str1)

	// 9.2 含有中文字符串
	fmt.Println("===含有中文字符串  []rune(str)===")
	str3 := "你好，世界！hello world！"
	s3 := []rune(str3) // 含有中文的字符串，需要用 `[]rune(str)` 处理
	s3[3] = '够'
	s3[4] = '浪'
	s3[12] = 'g'

	s3 = s3[:14]
	str3 = string(s3)
	fmt.Printf("str3 = %v\n", str3)
}

// 10. golang slice data[:6:8] 两个冒号的理解
func testMuli() {
	fmt.Println("===两个冒号的理解===")

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    d1 := slice[6:8]
	fmt.Printf("d1=%v，len(d1)=%v，cap(d1)=%v\n", d1, len(d1), cap(d1))

	// len = 6 - 0;  cap = 8 - 0
	d2 := slice[:6:8]
    fmt.Println(d2, len(d2), cap(d2))
	fmt.Printf("d2=%v，len(d2)=%v，cap(d2)=%v\n", d2, len(d2), cap(d2))
}

// 11.数组or切片转字符串：
func arrayOrSliceToString() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	str := strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", "~", -1)
	fmt.Printf("str = %v\n", str)

	// 关于 Replace 参数解释
	// 参考：https://www.cnblogs.com/haima/p/12697644.html
}


// 12.Slice底层实现
func sliceBase() {
	fmt.Println("===Slice底层实现===")
	// 关于切片和数组怎么选择？
	// 在 Go 中，与 C 数组变量隐式作为指针使用不同，Go 数组是值类型，赋值和函数传参操作都会复制整个数组数据。

	arrayA := [2]int{100, 200}
	var arrayB [2]int

	// Go数组是值类型，赋值操作会赋值整个数组的数据
	arrayB = arrayA

	// arrayA 与 arrayB 地址不同
	fmt.Printf("arrayA: %p，%v\n", &arrayA, arrayA)
	fmt.Printf("arrayB: %p，%v\n", &arrayB, arrayB)

	testArray(arrayA)


	/*
	打印结果：
		arrayA: 0xc0000c0650，[100 200]
		arrayB: 0xc0000c0660，[100 200]
		func Array: 0xc0000c0690 , [100 200]
	三个数组的内存地址都不同。这样就验证了 Go语言的数组赋值和函数传参都是值赋值的。这样导致什么问题❓❓❓

	问题：
	假想每次传参都用数组，那么每次数组都要被复制一遍。如果数组大小有 100万，在64位机器上就需要花费大约 800W 字节，即 8MB 内存。这样会消耗掉大量的内存。于是乎有人想到，函数传参用数组的指针。
	*/


	// 扩容是否生成新数组
	growSlice()
}

func testArray(x [2]int) {
	fmt.Printf("func Array: %p , %v\n", &x, x)
}

// 扩容是否生成新数组
func growSlice() {
	fmt.Println("===扩容是否生成新数组===")

	array := [4]int{10, 20, 30, 40}
    slice := array[0:2]
    newSlice := append(slice, 50)

    fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
    fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))

    newSlice[1] += 10

    fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
    fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
    fmt.Printf("After array = %v\n", array)

	/*
	Before slice = [10 20], Pointer = 0xc0000b4420, len = 2, cap = 4
	Before newSlice = [10 20 50], Pointer = 0xc0000b4438, len = 3, cap = 4

	After slice = [10 30], Pointer = 0xc0000b4420, len = 2, cap = 4
	After newSlice = [10 30 50], Pointer = 0xc0000b4438, len = 3, cap = 4

	After array = [10 30 50 40]		// 数组 array 的值也被影响了。
	*/




}
