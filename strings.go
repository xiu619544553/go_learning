package main

import (
	"fmt"
	"strings"
)

func aboutStrings() {
	fmt.Println("===字符串处理===")

	// 项目中会遇到切割需求
	strTest1()


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