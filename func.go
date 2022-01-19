package main

import (
	"fmt"
)

// 函数可以没有参数或接受多个参数。
// 注意类型在变量名之后 。
// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
// 函数可以返回任意数量的返回值。
// 使用关键字 func 定义函数，左大括号依旧不能另起一行。

func aboutFunction() {

	// 示例
	num, str := testFunc(1, 2, "abc")
	fmt.Printf("num=%v，str=%s\n", num, str)

	// 函数是第一类对象，可作为参数传递。建议将复杂签名定义为函数类型，以便于阅读
	s1 := testFunc1(func() int { return 100 }) // 直接将匿名函数当参数
	fmt.Printf("s1=%v\n", s1)

	s2 := format(func(s string, x, y int) string { 
		return fmt.Sprintf("%s_%d_%d", s, x, y)
	}, "abc", 100, 200)
	fmt.Printf("s2=%v\n", s2)


	// 引用传递
	var a, b int = 1, 2
	swap2(&a, &b)
	fmt.Printf("a=%d，b=%d\n", a, b)

	// 使用 slice 对象做变参时，必须展开。（slice...）
	s3 := []int{1, 2, 3}
	res := myFunc5("sum: %d", s3...)
	fmt.Printf("res = %s\n", res)
}

// x、y类型相同且连续，可以只保留最后一个变量的类型
func testFunc(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf("%s_%d", s, n)
}

/************************* start 函数作为参数**********************/
// 函数作为参数
func testFunc1(fn func() int) int {
	return fn()
}

// 定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

/************************* end 函数作为参数**********************/



/************************* start 参数**********************/

// 值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数
// 引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

// 值传递
func swap1(x, y int) {

}

// 引用传递
func swap2(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

// 在默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
// 注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
// 注意2：map、slice、chan、指针、interface默认以引用的方式传递。
// 不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）
// Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。
// 在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可。


// 注意：其中args是一个slice，我们可以通过arg[index]依次访问所有参数,通过len(arg)来判断传递参数的个数.
// 0或多个参数
func myFunc1(args ...int) {}

// 1或多个参数
func myFunc2(a int, args...int) {}

// 2或多个参数
func myFunc3(a int, b int, args...int) {}


// 任意类型的不定参数：就是函数的参数和每个参数的类型都不是固定的。
// 用interface{} 传递任意类型数据是Go语言的惯例用法，而且 interface{} 是类型安全的。

func myFunc4(args ...interface{}) {}


// 举例说明
func myFunc5(s string, n ...int) string {
	// n 是 slice
	var x int
	for _, v := range n {
		x += v
	}
	return fmt.Sprintf(s, x)
}

// 使用 slice 对象做变参时，必须展开。（slice...）



/************************* end 参数**********************/



/************************* start 返回值**********************/

// Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略。
// 多返回值可直接作为其他函数调用实参。


/************************* end 返回值**********************/

// 匿名函数
// https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%8C%BF%E5%90%8D%E5%87%BD%E6%95%B0.html