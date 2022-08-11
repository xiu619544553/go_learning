package main

import (
	"fmt"
	"reflect"
)

func aboutStructTag() {
	fmt.Println("Go语言结构体标签（Struct Tag）")

	// 通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的 Tag 被称为结构体标签（Struct Tag）。结构体标签是对结构体字段的额外信息标签。


	// 结构体标签格式
	// key1:"value1" key2:"value2"
	// 结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。


	// 从结构体标签中获取值

	// 根据 Tag 中的键获取对应的值
	// func(tag StructTag)Get(key string)string

	// 查询值是否存在
	// func(tag StructTag)Lookup(key string)(value string,ok bool)

	// 取值
	tagTest1()


	// 通过反射获取值信息
	tagTest2()
}

// 结构体标签使用
func tagTest1() {

	type tagcat struct {
		Name string
		Type int `json:"type" id:"100"`
	}

	typeOfCat := reflect.TypeOf(tagcat{})

	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"))
		fmt.Println(catType.Tag.Get("id"))
	}
}


// 通过反射获取值信息
func tagTest2() {

	// 使用反射值对象包装任意值
	// value := reflect.ValueOf(rawValue)
	

	// 从反射值对象获取被包装的值

}

