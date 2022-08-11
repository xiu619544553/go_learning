package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// 异常处理
func aboutError() {

	/*
	panic：
	    1、内置函数
		2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
		3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
		4、直到goroutine整个退出，并报告错误
	*/

	/*
	recover：
	    1、内置函数
		2、用来控制一个 goroutine 的 panicking 行为，捕获 panic，从而影响应用的行为
		3、一般的调用建议
			a). 在 defer 函数中，通过 recever 来终止一个 goroutine 的 panicking 过程，从而恢复正常代码的执行
			b). 可以获取通过 panic 传递的 error
	*/

	/*
	注意：
		1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
    	2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
    	3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。
	*/

	// errorTest1()

	// 向已关闭的通道发送数据，引发 panic
	// errorTest2()

	// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
	// errorTest3()

	// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。
	// errorTest4()

	// 使用延迟匿名函数或下面这样都是有效的。
	// errorTest5()

	// 如果需要保护代码段，可以将代码跨重构成匿名函数，如此可以保证后续代码被执行
	// errorTest6()

	// 除用 panic 引发中断性错误外，还可返回 error 类型错误对象来表示函数调用状态。
	// errorTest7()

	// Go实现类似 try catch 的异常处理
	// errorTest8()

	// 自定义error
	// errorTest9()

	// 自定义 error
	errorTest10()
}

func errorTest1() {
	defer func() {
		fmt.Println("o NO")
		if err := recover(); err != nil {
			fmt.Println(err.(string)) // 将 interface{} 转型为具体类型
		}
	}()
	
	panic("panic error!")	
}

// 向已关闭的通道发送数据，引发 panic
func errorTest2() {
	defer func ()  {
		if err := recover(); err != nil {
			// 输出结果：
			// send on closed channel
			fmt.Println(err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}

// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
func errorTest3() {
	defer func () {
		fmt.Println(recover()) // defer panic
	}()

	defer func () {
		panic("defer panic")
	}()

	panic("test panic")
}

// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。
func errorTest4() {
	defer func () {
		fmt.Println(recover()) // 有效：可以捕捉到 panic
	}()

	defer recover() // 无效：无法捕捉到 panic
	defer fmt.Println(recover()) // 无效：无法捕捉到 panic

	defer func ()  {
		func ()  {
			fmt.Println("defer inner")
			recover() // 无效，无法捕捉 panic
		}()
	}()

	panic("test panic")
}

// 使用延迟匿名函数或下面这样都是有效的。
func errorTest5() {
	defer errorTest5Except()
	panic("test5 panic")
}

func errorTest5Except() {
	fmt.Println(recover())
}

// 如果需要保护代码段，可以将代码跨重构成匿名函数，如此可以保证后续代码被执行
func errorTest6() {
	errorTest6IMP(2, 1)
}

func errorTest6IMP(x, y int) {
	var z int = 1

    func() {
        defer func() {
            if recover() != nil {
                z = 0
            }
        }()
        panic("test panic")
        z = x / y
        return
    }()

    fmt.Printf("x / y = %d\n", z)
}


// 除用 panic 引发中断性错误外，还可返回 error 类型错误对象来表示函数调用状态。

func errorTest7() {
	defer func ()  {
		fmt.Println(recover())
	}()

	switch z, err := errorTest7IMP(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}

var ErrDivByZero = errors.New("division by zero")

func errorTest7IMP(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x/y, nil
}

// Go实现类似 try catch 的异常处理

func errorTest8() {
	Try(func ()  {
		panic("模拟 panic")
	}, func (err interface{})  {
		fmt.Println(err)
	})
}

func Try(fun func(), handler func(interface{})) {
	defer func () {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}


/*===================================================*/
//                 系统抛异常和自己抛异常                 /
/*===================================================*/

// 系统抛异常和自己抛异常
func errorTest9() {

	// 捕获异常
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	
	// 分别注释下面的函数，defer 会捕捉到异常

	// 系统抛异常
	systemError()

	// 自己抛异常
	area := getCircleArea(-1)
	fmt.Printf("area：%v\n", area)
}

func systemError() {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)

	//a[10] = 11
	index := 10
	a[index] = 10
	fmt.Println(a)
}

func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
	   // 自己抛
	   panic("半径不能为负")
	}
	return 3.14 * radius * radius
 }


/*===================================================*/
//                    自定义 error                     /
/*===================================================*/

type PathError struct {
    path       string
    op         string
    createTime string
    message    string
}

func (p *PathError) Error() string {
    return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
        p.op, p.createTime, p.message)
}

func Open(filename string) error {

    file, err := os.Open(filename)
    if err != nil {
		// 此处为何可以返回 &PathError？？？？？？
        return &PathError{
            path:       filename,
            op:         "read",
            message:    err.Error(),
            createTime: fmt.Sprintf("%v", time.Now()),
        }
    }

    defer file.Close()
    return nil
}

 func errorTest10() {
	err := Open("/Users/hello/Desktop/testt.txt")
    switch v := err.(type) {
    case *PathError:
        fmt.Println("get path error,", v)
    default:
		fmt.Println("默认")
    }
 }