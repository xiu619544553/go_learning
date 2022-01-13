package main // 声明 main 包

import (
	"fmt" // 导入 fmt 包，打印字符串是需要用到
)

func main () {  // 声明 main 主函数
	fmt.Println("hello world!")

	// fmt.Println("==========基本数据类型==========")
	// baseDataType()

	// fmt.Println("==========数组 Array==========")
	// aboutArray()

	// fmt.Println("==========切片 Slice==========")
	// aboutSlice()

	fmt.Println("==========指针==========")
	aboutPointer()
}






