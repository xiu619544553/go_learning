package main

import (
	"fmt"
)

func aboutMap() {

	// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	// 定义：map[KeyType]ValueType
	//			keyType -- 键的类型
	// 			ValueType -- 值类型

	// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
	// make(map[KeyType]ValueType, [cap])
	// 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量


	// 1.先初始化，然后赋值
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 99
	scoreMap["李四"] = 100
	fmt.Printf("scoreMap = %v\n", scoreMap)				// scoreMap = map[张三:99 李四:100]
	fmt.Printf("scoreMap[张三]=%v\n", scoreMap["张三"])	 // scoreMap[张三]=99
	fmt.Printf("类型：%T\n", scoreMap)					 // 类型：map[string]int


	// 2.声明的时候赋值
	userInfo := map[string]string{
		"userName" : "Alex",
		"password" : "666",
	}
	fmt.Printf("userInfo = %v\n", userInfo) // userInfo = map[password:666 userName:Alex]


	// 3.判断某个键是否存在
}



