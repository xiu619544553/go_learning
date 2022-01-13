package main

import "fmt"



// for range
func aboutForRange() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for i, v := range slice {
		// ❌这是错误的
		// myMap[i] = &v

		// 创建一个新的变量，如此方能保证value对应的地址不同
		num := v
		myMap[i] = &num
	}

	fmt.Println("===new map===")
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
		// 输出结果如下，不是预期结果
		// map[1]=3
		// map[2]=3
		// map[3]=3
		// map[0]=3

		// 是什么原因导致映射所有值都相同
		// 原因分析：因为for range创建了每个元素的副本，而不是直接返回每个元素的引用，
		// 如果使用该值变量的地址作为指向每个元素的指针，就会导致错误，在迭代时，返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以值的地址总是相同的，导致结果不如预期。
	}
}