package main

import (
	"fmt"
	"errors"
	"net/http"
)

func aboutDefer() {
	fmt.Println("defer")
	// https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%BB%B6%E8%BF%9F%E8%B0%83%E7%94%A8defer.html
	/*
	defer特性：
		1. 关键字 defer 用于注册延迟调用。
		2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
		3. 多个defer语句，按先进后出的方式执行。
		4. defer语句中的变量，在defer声明时就决定了。

	defer用途：
	    1. 关闭文件句柄
    	2. 锁资源释放
    	3. 数据库连接释放
	*/

	defer1()

	// defer 和闭包
	defer2()

	// 易犯错：搭配 for-range 使用，容易出问题的地方
	defer3()
	defer4()
	defer5()

	// 多个 defer
	defer6()

	// 延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。
	defer7()

	// 滥用 defer 可能会导致性能问题，尤其是在一个 "大循环" 里。
	defer8()

	// defer陷阱
	defer9()

	// defer 与 return
	defer10()

	// defer nil 函数
	// defer11()

	// 在正确的位置使用 defer
	defer12()
}

// 示例1
func defer1() {
	var arr [5]int

	for i, _ := range arr {
		defer fmt.Println(i)
	}
	/*
	输出结果：（先进后出）
		4
		3
		2
		1
		0
	*/
}

// 示例2：defer 和闭包
func defer2() {
	var arr [5]int 
	for i, _ := range arr {
		// 每次“defer”语句执行时，调用的函数值和参数会像往常一样计算并重新保存，但实际的函数不会被调用。
		// 也就是说函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.
		defer func () {
			fmt.Println(i)
		}()
	}

	/*
	输出结果：
		4
		4
		4
		4
		4
	*/ 
}


// 示例3❌
func defer3() {

	fmt.Println("===defer3===")

	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}

	for _, v := range ts {
		/*
			closed...t地址：0xc00010a300，name=c
			closed...t地址：0xc00010a300，name=c
			closed...t地址：0xc00010a300，name=c
		注意：这是 for.range 的特殊之处，v 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值。
			 defer 延迟加载，因为先进后出，最后进入的 c.name=c 会覆盖前面的值。
			 defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。
		*/
		fmt.Printf("v地址:：%p，name=%s\n", &v, v.name)
		defer v.Close()
	}
}

type Test struct {
	name string
}

func (t *Test) Close() {
	// fmt.Println(t.name, "closed...")
	fmt.Printf("closed...t地址：%p，name=%s\n", t, t.name)
}

// 示例4✅
func defer4() {

	fmt.Println("===defer4===")

	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}
	
	for _, v := range ts {
		defer Close(v)
	}
}

// 多声明了一个函数
func Close(t Test) {
	t.Close()
}

// 当然,如果你不想多写一个函数,也很简单,可以像下面这样
// 示例5✅
func defer5() {
	fmt.Println("===defer5===")

	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}
	
	for _, v := range ts {
		tmp := v
		defer tmp.Close()
	}
}


// 多个 defer 
// 多个 defer 注册，按 FILO 次序执行 ( 先进后出 )。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。
func defer6() {
	fmt.Println("=== 多个 defer ===")
	defer6Test(0)
}

func defer6Test(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")

	defer func() {
		// println(100/x) // 哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。
	}()

	defer fmt.Println("c")
}

// 延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。
func defer7() {
	fmt.Println("===延迟调用参数在注册时求值或复制，可用指针或闭包 延迟 读取。===")
	x, y := 10, 20

	defer func(i int) {
		fmt.Printf("defer：i=%d，y=%d\n", x, y) // y 闭包引用。读取到的是 `y += 100` 表达式计算后的值
	}(x) // x 被复制，在注册时复制

	x += 10
	y += 100
	fmt.Printf("x=%d，y=%d\n", x, y)
}

// 滥用 defer 可能会导致性能问题，尤其是在一个 "大循环" 里。
func defer8() {
	// https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%BB%B6%E8%BF%9F%E8%B0%83%E7%94%A8defer.html
}


// defer陷阱
func defer9() {
	fmt.Println("===defer陷阱===")
	foo(2, 0)
}

func foo(a, b int) (i int, err error) {

	/*
	输出结果：
		third defer err divided by zero!
		second defer err <nil>
		first defer err <nil> 

	解释：如果 defer 后面跟的不是一个 closure 最后执行的时候我们得到的并不是最新的值。
	*/

	// ①
    defer fmt.Printf("first defer err %v\n", err)

	// ②匿名函数
    defer func(err error) {
		fmt.Printf("second defer err %v\n", err) 
	}(err)
		
	// ③闭包
    defer func() {
		fmt.Printf("third defer err %v\n", err) 
	}()
	
    if b == 0 {
        err = errors.New("divided by zero!")
        return
    }

    i = a / b
    return
}


// defer 与 return
func defer10() {
	fmt.Println("===defer 与 return===")
	foo1()
}

func foo1() (i int) {
	i = 0
	defer func ()  {
		fmt.Println(i)
	}()
	
	return 2
}


// defer nil 函数
/*
输出日志：
runs
runtime error: invalid memory address or nil pointer dereference
*/
func defer11() {
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	test11()
}

func test11() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

// 在正确的位置使用 defer
func defer12() {
	fmt.Println("===在正确的位置使用 defer===")
	fmt.Printf("error = %v\n", defer12Test())
}

func defer12Test() error {
	res, err := http.Get("https://www.baidu.com")

	if res != nil { // 必须判空，否则当 res==nil 时， defer 中使用 res 会crash
		defer res.Body.Close()
	}

	if err != nil {
		return err
	}

	return nil
}

