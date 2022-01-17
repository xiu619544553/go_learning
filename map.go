package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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
	initMap()

	// 2.声明的时候赋值
	statementMap()

	// 3.判断某个键是否存在
	// 4.遍历map
	// 5.删除键值对
	judgmentMap()

	// 6.按照指定顺序排序
	sortMap()

	// 7.元素为 map 类型的切片
	mapSlice()

	// 8.值为切片类型的map
	mapValueSlice()

	// 9.Go中Map的使用
	goMapUsage()
}

// 1.先初始化，然后赋值
func initMap() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 99
	scoreMap["李四"] = 100
	fmt.Printf("scoreMap = %v\n", scoreMap)				// scoreMap = map[张三:99 李四:100]
	fmt.Printf("scoreMap[张三]=%v\n", scoreMap["张三"])	 // scoreMap[张三]=99
	fmt.Printf("类型：%T\n", scoreMap)					 // 类型：map[string]int
}

// 2.声明的时候赋值
func statementMap() {
	userInfo := map[string]string{
		"userName" : "Alex",
		"password" : "666",
	}
	fmt.Printf("userInfo = %v\n", userInfo) // userInfo = map[password:666 userName:Alex]
}

// 3.判断某个键是否存在
// 4.遍历map
// 5.删除键值对
func judgmentMap() {
	// 格式
	// value, ok := map[key]

	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	scoreMap["马武"] = 110

	// 如果key存在，ok为true，v为对应的值；不存在 ok为false，v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Printf("v = %v\n", v)
	} else {
		fmt.Println("查无此人")
	}

	// 4.遍历map
	// 注意：遍历map时的元素顺序与添加键值对的顺序无关。

	// 遍历 key、value
	traversalMap(scoreMap)

	// 遍历 key
	for k := range scoreMap {
		fmt.Printf("k=%v\n", k)
	}

	// 5.删除键值对
	// delete(map, key)
	delete(scoreMap, "张三")
	traversalMap(scoreMap)
}

// 遍历 key、value
func traversalMap(x map[string]int) {
	fmt.Println("======遍历map=====")
	for k, v := range x {
		fmt.Printf("k=%v，v=%v\n", k, v)
	}
}

// 6.按照指定顺序排序
func sortMap() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	fmt.Println("===按照指定顺序排序===")
	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	// 取出 map 中所有key存入切片 keys
	var keys = make([]string, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// 对切片排序
	sort.Strings(keys)

	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Printf("key=%v，scoreMap[%v]=%v\n", key, key, scoreMap[key])
	}
}

// 7.元素为 map 类型的切片
func mapSlice() {
	fmt.Println("===元素为 map 类型的切片===")
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// 对切片中的 map进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "张三"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"]  = "首都北京"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// 8.值为切片类型的map
func mapValueSlice() {
	fmt.Println("===值为切片类型的map===")
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")

	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Printf("sliceMap = %v\n", sliceMap)
}

// 9.Go中Map的使用
func goMapUsage() {
	fmt.Println("===9.Go中Map的使用===")
	//直接创建初始化一个map
	var mapInit = map[string]string {"xiaoli":"湖南", "xiaoliu":"天津"}
	fmt.Printf("mapInit=%v\n", mapInit)

	//声明一个map类型变量,
	//map的key的类型是string，value的类型是string
	var mapTemp map[string]string
	//使用make函数初始化这个变量,并指定大小(也可以不指定)
	mapTemp = make(map[string]string,10)
	//存储key ，value
	mapTemp["xiaoming"] = "北京"
	mapTemp["xiaowang"]= "河北"
	//根据key获取value,
	//如果key存在，则ok是true，否则是flase
	//v1用来接收key对应的value,当ok是false时，v1是nil
	v1,ok := mapTemp["xiaoming"]
	fmt.Println(ok,v1)
	//当key=xiaowang存在时打印value
	if v2,ok := mapTemp["xiaowang"]; ok{
		fmt.Println(v2)
	}
	//遍历map,打印key和value
	for k,v := range mapTemp{
		fmt.Println(k,v)
	}
	//删除map中的key
	delete(mapTemp,"xiaoming")
	//获取map的大小
	l := len(mapTemp)
	fmt.Println(l)
}