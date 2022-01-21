package main

import (
	"fmt"
)

func aboutFlow() {

	// switch
	aboutSwitch()

	// for 
	aboutFor()

	// range
	aboutRange()

	// goto 
	aboutGoto()

	// continue
	aboutContinue()

	// break
	aboutBreak()
}

// switch
func aboutSwitch() {
	// 用法1
	var score int = 90
	var level string = "B"
	switch score {
	case 90: 
		level = "A"
	case 80: 
		level = "B"
	case 70: 
		level = "C"
	default: 
		level = "D"
	}
	fmt.Println(level)


	// Type Switch
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T\n", i)
	case int:
		fmt.Printf("x 是 int 类型")
	case float64:
		fmt.Printf("x 是 float64 类型")
	default:
		fmt.Printf("未知型\n")
		
	}


	// falltrhough
	var k = 0
	switch k {
	case 0:
		fmt.Printf("fallthrough：%v\n", k)
		fallthrough // 使用fallthrough强制执行后面的case代码
	case 1:
		fmt.Printf("1：%v\n", k)
	default:
		fmt.Printf("default：%v\n", k)
	}
}

// select
func aboutSelect() {
	// select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
}

// for 
func aboutFor() {
	s := "abc"

	// for 
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%v] = %c\n", i, s[i]) // %c 输出字符
	}

	// for range
	for index, value := range s {
		fmt.Printf("s[%v] = %c\n", index, value) 
	}

	// for -- 替代 while 
	n := len(s)
	for n > 0 {
		n--
		fmt.Printf("s[%v] = %c\n", n, s[n])
	}
}

// range
// range类似迭代器操作，返回 (索引, 值) 或 (键, 值)。
func aboutRange() {
	fmt.Println("===range===")

	/*
	for range 可以遍历数组、切片、字符串、map 及通道（channel）。一般形式如下：

	for key, val := range coll {
    	...
	}

	注意：val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值。一个循环中，val的地址不会发生变化。
	*/ 


	// range 会复制对象
	// 验证
	a := [3]int{1, 2, 3}

	fmt.Printf("a[0]地址：%p\n", &a[0])
	fmt.Printf("a[1]地址：%p\n", &a[1])
	fmt.Printf("a[2]地址：%p\n", &a[2])

	/*
	输出日志：
		a[0]地址：0xc0000140c0
		a[1]地址：0xc0000140c8
		a[2]地址：0xc0000140d0
		index=0，value=1，index地址：0xc000016078，value地址：0xc0000160b0
		index=1，value=2，index地址：0xc000016078，value地址：0xc0000160b0
		index=2，value=3，index地址：0xc000016078，value地址：0xc0000160b0

	结论：for.range，会拷贝一份 a，index、value是新的变量，接收备份中的值
	*/

	for index, value := range a { // index、value都是从复制品中取出

		fmt.Printf("index=%d，value=%d，index地址：%p，value地址：%p\n", index, value, &index, &value)

		if index == 0 {
			a[1], a[2] = 999, 999
			fmt.Printf("确认修改是否有效：%v\n", a) // 确认修改是否有效：[1 999 999]
		}

		// index、value都是从复制品中取出。
		/*
		index=0，value=1
		index=1，value=2
		index=2，value=3

		结论：range 复制了一份 a，index、value都是从复制品中取出的。不受 a[1]、a[2]修改值的影响
		*/ 
		fmt.Printf("index=%v，value=%v\n", index, value)

		a[index] = value + 100 // 使用复制品中取出的 value 修改原数组
	}
	
	fmt.Printf("a = %v\n", a) // a = [101 102 103]


	// 推荐使用引用类型：其底层数据不会被复制
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("s地址：%p\n", &s)
	fmt.Printf("s[0]地址：%p\n", &s[0])
	fmt.Printf("s[1]地址：%p\n", &s[1])
	fmt.Printf("s[2]地址：%p\n", &s[2])
	fmt.Printf("s[3]地址：%p\n", &s[3])
	fmt.Printf("s[4]地址：%p\n", &s[4])

	/*
	输出内容：
		s[0]地址：0xc00012a000
		s[1]地址：0xc00012a008
		s[2]地址：0xc00012a010
		s[3]地址：0xc00012a018
		s[4]地址：0xc00012a020
		index=0，value=1，index地址：0xc000114008，value地址：0xc000114040
		index=1，value=2，index地址：0xc000114008，value地址：0xc000114040
		index=2，value=3，index地址：0xc000114008，value地址：0xc000114040
		index=3，value=4，index地址：0xc000114008，value地址：0xc000114040
		index=4，value=5，index地址：0xc000114008，value地址：0xc000114040

	*/

	for index, value := range s {
		
		// fmt.Printf("index=%d，value=%d，index地址：%p，value地址：%p\n", index, value, &index, &value)

		if index == 0 {
			s = s[:3] 				// 对 slice 的修改，不会影响 range。
			fmt.Printf("s=%v\n", s) // s=[1 2 3]
			s[2] = 100 				// 对底层数据的修改。
			fmt.Printf("s地址：%p\n", &s)
		}

		fmt.Printf("index=%v，value=%v\n", index, value)
	}

	fmt.Printf("s====%v\n", s) // s====[1 2 100]

}

/************* goto **** start *******/  
func aboutGoto() {

	// 下面这段代码在满足条件时，需要连续退出两层循环，使用传统的编码方式如下：
	aboutGeneralPractice()

	// 使用 goto 优化上述代码
	aboutOptimizeWithGoto()

	// 使用 goto 集中处理错误
	// 参考：http://c.biancheng.net/view/49.html
}

func aboutGeneralPractice() {
	var breakAgain bool
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				// 设置标记
				breakAgain = true
				// 退出本次循环
				break
			}
		}
		// 根据标记，退出循环
		if breakAgain == true {
			break
		}
	}
}

func aboutOptimizeWithGoto() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				// 跳转到标签
				goto breakHere
			}
		}
	}

	// 手动返回, 避免执行进入标签
    return

	// 标签
breakHere:
	fmt.Println("我是标签")	
}
/************* goto **** end *******/  



/************* continue **** start *******/  
func aboutContinue() {

	// 用法1：Go语言中 continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用
	aboutContinueUsage1()

	// 用法2：配合标签使用
}

// Go语言中 continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用
func aboutContinueUsage1() {
	var count int
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		count ++
	}
	fmt.Printf("记录循环次数：%v\n", count)
}

// 配合标签使用：在 continue 语句后添加标签时，表示开始标签对应的循环
func aboutContinueUsage2() {

	var flag int = 0

	// 标签
OuterLoop:	
	for i := 0; i < 5; i++ {
		flag ++
		if flag == 2 {
			// 跳转到标签
			continue OuterLoop	
		}
	}		

}

/************* continue **** end *******/


/************* break **** start *******/

func aboutBreak() {
	// 用法1：结束for、switch 和 select 的代码块
	aboutBreakUsage1()

	// 用法2： break 语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的 for、switch 和 select 的代码块上。
	aboutBreakUsage2()
}

// 用法1
func aboutBreakUsage1() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			// 结束 for 循环
			break
		}
	}
}

// 用法2
func aboutBreakUsage2() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}

/************* break **** end *******/