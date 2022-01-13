package main

import (
	"fmt"
)

func aboutPointer() {
	fmt.Println("===指针===")

	/*
	区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。
	要搞明白Go语言中的指针需要先知道3个概念：
		①指针地址、
		②指针类型
		③指针取值。



	1.1 Go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。
		传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。Go语言中的指针操作非常简单，只需要记住两个符号：&（取地址）和*（根据地址取值）。

	取变量指针的语法：prt := &v
		* v:代表被取地址的变量，类型为T
		* ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针
	*/

	test1()
	testNew()
	testMake()
}

func test1() {
	a := 10
	b := &a
	fmt.Printf("a: %d  ptr: %p\n", a, &a)
	fmt.Printf("b: %p  type: %T\n", b, b)  // %T：输出变量的类型和值
	fmt.Println(&b)
}

 // new
 func testNew() {

	// var a *int
	// *a = 100
	// fmt.Printf("*a = %d\n", *a)
	// ❌报错panic：
	// 原因分析：在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。

	// new 函数
	// 定义：func(Type) *Type
	// 不太常用

	a := new(int)
	b := new(bool)

	fmt.Printf("a类型是：%T\n", a) // a类型是：*int
	fmt.Printf("b类型是：%T\n", b) // b类型是：*bool

	fmt.Printf("a的值：%v\n", *a) // a的值：0
	fmt.Printf("b的值：%v\n", *b) // b的值：false

	// 修复上面❌报错代码，✅正确写法
	var c *int = new(int)
	*c = 10
	fmt.Printf("*c = %d\n", *c) // *c = 10
 }

 // make
 // make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
func testMake() {
	// make(type, 0)
	// func(t Type, size ...IntegerType) Type
	// make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。
	
	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}
 
// 1.二者都是用来做内存分配的。
// 2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
// 3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。