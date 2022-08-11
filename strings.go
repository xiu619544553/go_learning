package main

import (
	"fmt"
	"strings"
)

func aboutStrings() {
	fmt.Println("===字符串处理===")

	// 项目中会遇到切割需求
	strTest1()

	// 字符串操作
	strTest2()
}

// 项目中会遇到切割需求
func strTest1() {
	// 问题：重复代码
	// @第一个字符
	idx1 := strings.Index("alex@163.com", "@")
	fmt.Println("idx1 =", idx1)

	idx2 := strings.Index("alex@@163.com", "@")
	fmt.Println("idx2 =", idx2)

	// @最后一个字符
	idx3 := strings.LastIndex("alex@@163.com", "@")
	fmt.Println("idx3 =", idx3)
}

func strTest2() {
	//（1）字符串是常量，可以通过类似数组 的索引访问其字节单元，但是不能修改某个字节的 值。例如 :
	var a = "hello world"
	b := a[1]
	fmt.Println("b =", b)

	//（2）宇符串转换为切片 []byte(s)要慎用，尤其是当数据量较大时(每转换一次都需复制内容)。例如:
	c := "hello world"
	d := []byte(c)
	fmt.Println("d =", d) // d = [104 101 108 108 111 32 119 111 114 108 100]
	fmt.Printf("%T", d)   // []uint8

	//（3）字符串尾部不包含 NULL 字符，这一点和 C/C++不一样。
	//（4）字符串类型底层实现是一个二元的数据结构，一个是指针指向字节数组的起点，另一个是长度 。 例如
	// runtime/string .go
	// type stringStruct struct {
	// 	str unsafe.Pointer  // 指向底层字节数组的指针
	// 	len int				// 字节数组长度
	// }

	//（5）基于字符串创建的切片和原字符串指向相同的底层字符数组 ， 串的切片操作返回的子串仍然是由ing，而非 slic巳。 例如：
	e := a[0:4]
	f := a[1:]
	g := a[:4]
	fmt.Printf("e=%v，f=%v，g=%v\n", e, f, g)

	//（6）字符串和切片的转换:字符串可以转换为字节数组，也可以转换为 Unicode 的字数组。 例如 :
	h := "helllo，世界！"
	h1 := []byte(a)
	h2 := []rune(a)
	fmt.Printf("（6）h=%v，h1=%v，h2=%v\n", h, h1, h2)

	//（7）字符串的运算。例如：
	i := "hello"
	j := "world"
	k := i + j // 字符串拼接
	fmt.Printf("k.len = %v\n", len(k)) // 内置函数 len，获取字符串长度

	// 遍历字节数组
	for i := 0; i < len(k); i++ {
		fmt.Printf("k[%v] = %v\n", i, k[i])
	}

	// 遍历rune数组
	for i, v := range k {
		fmt.Printf("i = %v，v = %v\n", i, v)
	}
}