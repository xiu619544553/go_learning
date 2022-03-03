package main

import (
	"fmt"
	"reflect"
)

func aboutReflect() {
	fmt.Println("反射")
	
	// 反射的概念
	// Go语言提供了一种机制在运行时更新和检查变量的值、调用变量的方法和变量支持的内在操作，但是在编译时并不知道这些变量的具体类型，这种机制被称为反射。
	// 反射也可以让我们将类型本身作为第一类的值类型处理。

	// 反射是指在程序运行期对程序本身进行访问和修改的能力。
	// 程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

	// 支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。


	// reflect包
	reflectTest1()


	// 理解反射，首先理解type和kind的区别。
	// 编程中，type用的最多。但在反射中，当需要区分一个大品种的类型，就会用到种类kind。
	// 比如，需要统一判断类型中的指针，使用kind就较为方便。


	// 1、反射 kind 的定义    http://c.biancheng.net/view/4407.html
	// 种类指的是对象归属的品种
	// type Kind uint
	// const (
	// 	Invalid Kind = iota  // 非法类型
	// 	Bool                 // 布尔型
	// 	Int                  // 有符号整型
	// 	Int8                 // 有符号8位整型
	// 	Int16                // 有符号16位整型
	// 	Int32                // 有符号32位整型
	// 	Int64                // 有符号64位整型
	// 	Uint                 // 无符号整型
	// 	Uint8                // 无符号8位整型
	// 	Uint16               // 无符号16位整型
	// 	Uint32               // 无符号32位整型
	// 	Uint64               // 无符号64位整型
	// 	Uintptr              // 指针
	// 	Float32              // 单精度浮点数
	// 	Float64              // 双精度浮点数
	// 	Complex64            // 64位复数类型
	// 	Complex128           // 128位复数类型
	// 	Array                // 数组
	// 	Chan                 // 通道
	// 	Func                 // 函数
	// 	Interface            // 接口
	// 	Map                  // 映射
	// 	Ptr                  // 指针
	// 	Slice                // 切片
	// 	String               // 字符串
	// 	Struct               // 结构体
	// 	UnsafePointer        // 底层指针
	// )

	// Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。


	// 2、从类型对象中获取类型名称和种类
	

	// 指针与指针指向的元素
	// Go语言程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型，这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作，代码如下：
	reflectTest2()


	// 使用反射获取结构体的成员类型
	// 任意值通过 reflect.TypeOf() 获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 reflect.Type 的 NumField() 和 Field() 方法获得结构体成员的详细信息。
	reflectTest3()


	// 结构体标签（Struct Tag）
	// 通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的 Tag 被称为结构体标签（StructTag）。结构体标签是对结构体字段的额外信息标签。
	// JSON、BSON 等格式进行序列化及对象关系映射（Object Relational Mapping，简称 ORM）系统都会用到结构体标签，这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。
	// 这些信息都是静态的，无须实例化结构体，可以通过反射获取到。

	// 1) 结构体标签的格式
	// tag在结构体字段后书写的格式如下：
	// `key1:"value1" key2:"value2"`
	// 书写规则如下：结构体标签由一对或者多对键值对组成；键与值使用冒号分隔，键与值之间不能有空格，值用双引号括起来；键值对之间使用一个空格分割；

	// 2) 从结构体标签中获取值
	// StructTag 拥有一些方法，可以进行 Tag 信息的解析和提取，如下所示：
	// func (tag StructTag) Get(key string) string：根据 Tag 中的键获取对应的值，例如`key1:"value1" key2:"value2"`的 Tag 中，可以传入“key1”获得“value1”。
	// func (tag StructTag) Lookup(key string) (value string, ok bool)：根据 Tag 中的键，查询值是否存在。

	// 3) 结构体标签格式错误导致的问题
	// 编写 Tag 时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误
	reflectTest4()



	// Go语言中的类型
	// Go语言是一门静态类型的语言，每个变量都有一个静态类型，类型在编译的时候确定下来。
	// 接口是一个重要的类型，它意味着一个确定的方法集合，一个接口变量可以存储任何实现了接口的方法的具体值（除了接口本身）
	// 接口类型中有一个极为重要的例子：空接口 
	// interface{}
	// 它表示了一个空的方法集，一切值都可以满足它，因为它们都有零值或方法。
	// go语言的接口是静态类型的。

	// 反射第一定律：反射可以将“接口类型变量”转换为“反射类型对象”
	// 注：这里反射类型指 reflect.Type 和 reflect.Value。
	// 
	reflectTest5()


	// 反射第三定律：如果要修改“反射类型对象”其值必须是“可写的”
	reflectTest6()


	// 结构体--反射
	reflectTest7()


	// reflect.Elem()
	// 通过反射获取指针指向的元素类型
	reflectTest8()
}

type MyStruct struct {}

// 反射的类型对象 reflect.Type
func reflectTest1() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println("typeOfA.Name() = ", typeOfA.Name())
	fmt.Println("typeOfA.Kind() = ", typeOfA.Kind())

	st := MyStruct {}
	typeOfS := reflect.TypeOf(st)
	fmt.Println("typeOfS.Name() = ", typeOfS.Name())
	fmt.Println("typeOfS.Kind() = ", typeOfS.Kind())

	st1 := &MyStruct{}
	typeOfS1 := reflect.TypeOf(st1)
	fmt.Println("typeOfS1.Name() = ", typeOfS1.Name())
	fmt.Println("typeOfS1.Kind() = ", typeOfS1.Kind())
}

// 指针与指针指向的元素
func reflectTest2() {
	fmt.Println("===指针与指针指向的元素===")
	type MyCat struct {}

	// 创建实例
	ins := &MyCat{}
	
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	// 显示反射类型对象的名称和种类
    fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())

	// 取类型的元素
	typeOfElem := typeOfCat.Elem()

	// 显示反射类型对象的名称和种类
    fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfElem.Name(), typeOfElem.Kind())
}

// 使用反射获取结构体的成员类型
func reflectTest3() {
	fmt.Println("=== 使用反射获取结构体的成员类型 ===")
	type MyAnimals struct {
		age int
		name string
	}

	// 创建实例
	an := MyAnimals{
		20,
		"Dog",
	}

	// 获取反射类型对象
	typeOfAn := reflect.TypeOf(an)


	fmt.Println("typeOfAn.Field(0) = ", typeOfAn.Field(0))

	for i := 0; i < typeOfAn.NumField(); i++ {

		// 获取结构体每个成员变量字段
		fieldType := typeOfAn.Field(i)

		fmt.Printf("name: %v， tag: %v\n", fieldType.Name, fieldType.Tag)
	}
}

// 结构体标签（Struct Tag）
func reflectTest4() {
	fmt.Println("=== 结构体标签（Struct Tag） ===")

	type Pig struct {
		weight float32
		Name string `json:"name" id:"66"`
	}

	typeOfCat := reflect.TypeOf(Pig{})
	if pigType, ok := typeOfCat.FieldByName("Name"); ok {
		fmt.Println(pigType.Tag.Get("json"))
	}
}

// Go语言中的类型
func reflectTest5() {

	fmt.Println("===Go语言中的类型===")	

	var x float64 = 3.4

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("reflect.TypeOf(x) = ", t)
	fmt.Println("reflect.ValueOf(x) = ", v)

	fmt.Println("type：", v.Type())
	fmt.Println("kind is float64：", v.Kind() == reflect.Float64)
	fmt.Println("value：", v.Float())

	// reflect.TypeOf(x) =  float64
	// reflect.ValueOf(x) =  3.4
}

// 反射第三定律：如果要修改“反射类型对象”其值必须是“可写的”
func reflectTest6() {
	fmt.Println("=== 反射第三定律：如果要修改“反射类型对象”其值必须是“可写的” ===")


	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// v.SetFloat(7.1)


	// 上述代码报错
	// "reflect: reflect.Value.SetFloat using unaddressable value"
	// 这里问题不在于值 7.1 不能被寻址，而是因为变量 v 是`不可写`的，`可写性`是反射类型变量的一个属性，但不是所有的反射类型变量都拥有这个属性
	// 检测 reflect.Value 变量的 `可写性`，对于上述的例子，可以这样写：
	var x float64 = 3.4
	v := reflect.ValueOf(x) // 此处传递的x是一个x的拷贝，并非x自身。所以 v.SetFloat(7.1) 并不能修改 x 的值，为了避免迷惑开发者，所以此处会报错 
	fmt.Println("settability of v:", v.CanSet())


	// 什么是`可写性`？
	// 类似于寻址能力，但是更严格，它是反射类型变量的一种属性，赋予该变量修改底层存储数据的能力。
	// “可写性”最终是由一个反射对象是否存储了原始值而决定的。
	// 反射的工作机制与此相同，如果想通过反射修改变量 x，就要把想要修改的变量的指针传递给反射库。


	var y float64 = 3.4
	z := reflect.ValueOf(&y)
	fmt.Println("type of z:", z.Type())
	fmt.Println("settability of z:", z.CanSet())  // settability of z: false


	// 反射对象 p 是不可写的，但是我们也不像修改 p，事实上我们要修改的是 *p。
	// 为了得到 p 指向的数据，可以调用 Value 类型的 Elem 方法。Elem 方法能够对指针进行“解引用”，然后将结果存储到反射 Value 类型对象 v 中：
	var a float64 = 3.4
	b := reflect.ValueOf(&a)
	c := b.Elem()
	fmt.Println("settability of c:", c.CanSet()) // settability of c: true

	// 只需要记住：只要反射对象要修改它们表示的对象，就必须获取它们表示的对象的地址。
}

// 结构体--反射
func reflectTest7() {

	fmt.Println("===结构体--反射===")

	// 我们一般使用反射修改结构体的字段，只要有结构体的指针，我们就可以修改它的字段。

	// 用结构体的地址创建反射变量，再修改它。然后我们对它的类型设置了 typeOfT，并用调用简单的方法迭代字段

	// T 字段名之所以大写，是因为结构体中只有可导出的字段是“可设置”的。
	type RT struct {
		A int
		B string
	}

	rt := RT{23, "Skip"}
	s  := reflect.ValueOf(&rt).Elem() // Elem()  It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s=%v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}


	// 因为 s 包含了一个可设置的反射对象，我们可以修改结构体字段
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset strip")
	fmt.Println("rt is now", rt)
}

// reflect.Elem()
// 通过反射获取指针指向的元素类型
func reflectTest8() {
	fmt.Println("=== 通过反射获取指针指向的元素类型 ===")

	// Go语言程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型。这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作，代码如下：

	type rcat struct {}

	// 创建cat实例
	ins := &rcat{}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n",typeOfCat.Name(), typeOfCat.Kind())

	// 取类型的元素
	elem := typeOfCat.Elem()

	// 显示反射类型对象的名称和种类
    fmt.Printf("element name: '%v', element kind: '%v'\n", elem.Name(), elem.Kind())
}