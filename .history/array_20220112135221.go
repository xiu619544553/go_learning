package main

import (
	"fmt"
)

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
	var arr5 = [...]struct { // 结构体数组
		name string
		age uint8
	}{
		{"user1", 10}, // 可省略元素类型
		{"user2", 20}, // 别忘了最后一行的逗号
	}
	fmt.Println("结构体数组：", arr5)
	

	// 比较两个数组是否相等。（包括数组的长度，数组中元素的类型）
	// 不能比较两个类型不同（元素个数不同）的数组，否则程序将无法完成编译。
	arr7 := [2]int{1, 2}
	arr8 := [2]int{1, 2}
	if arr7 == arr8 {
		fmt.Printf("arr7 == arr8")
	}

	
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

	// 内置函数 len、cap 都返回数组长度（元素数量）
	numsCount := cap(nums)
	numsCount2 := len(nums)
	fmt.Printf("numsCount=%d\n", numsCount)
	fmt.Printf("numsCount2=%d\n", numsCount2)

	// 指针数组 [n]*T，数组指针 *[n]T。
	// 数组指针：首先是一个指针，一个数组的地址
	// 指针数组：首先是一个数组，存储的数据类型是指针
	// 参考：https://blog.csdn.net/weixin_40123451/article/details/122092190



	// 多维数组
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
    b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
    fmt.Println(a, b)


	// 遍历多维数组
	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

    for k1, v1 := range f {
        for k2, v2 := range v1 {
            fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
        }
        fmt.Println()
    }


	// 注意：值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。

	fmt.Println("===验证 数组 是值类型===")
	arr6 := [2]int{}
	fmt.Printf("a: %p\n", &arr6)

	test(arr6)
	fmt.Println(arr6)

	
	fmt.Println("===数组拷贝和传参===")
	// 数组拷贝和传参
	var arr1 [5]int
    printArr(&arr1)
    fmt.Println(arr1)

    arr2 := [...]int{2, 4, 6, 8, 10}
    printArr(&arr2)
    fmt.Println(arr2)

	
	fmt.Println("===求元素的和===")
}

// 验证 数组 是值类型
func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 100
}

// 
func printArr(arr *[5]int) {
    arr[0] = 10
    for i, v := range arr {
        fmt.Println(i, v)
    }
}