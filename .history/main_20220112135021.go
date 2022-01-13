package main // 声明 main 包

import (
	"fmt" // 导入 fmt 包，打印字符串是需要用到
	"strings"
)

func main () {  // 声明 main 主函数
	fmt.Println("hello world!")

	

	fmt.Println("==========Array==========")
	aboutArray()
}

func getData() (int, int) {
	return 100, 200
}

func traversalString() {
	s := "Ha你好"
	
	count := len(s)
	fmt.Println("count =", count)
	
	for i := 0; i < len(s); i++ {
		// uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
		fmt.Printf("%d：%v(%c)", i, s[i], s[i])
		fmt.Println()
	}

	for index, r := range s {
		// rune
		fmt.Printf("%v，%v(%c)\n", index, r, r)
	}
	
}




