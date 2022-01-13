package main

import (
	"fmt"
)

func aboutSlice() {
	/*
	需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

	1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。 
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string 、 var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
	*/

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

	fmt.Println("===append 前=")
	fmt.Printf("data = %v\n", data)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("s.len = %v，cap = %v\n", len(s), cap(s)) // s.len = 2，cap = 3

	s = append(s, 100, 200)  // 一次append两个元素，超出了 s.cap 限制
	fmt.Printf("data = %v\n", data)
	fmt.Printf("s = %v\n", s)

	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
}