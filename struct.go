package main

import (
	"encoding/json"
	"fmt"
)

// 类型定义
type NewInt int

// 类型别名
type MyInt = int


func aboutStruct() {
	fmt.Println("===struct===")

	// 1.类型别名和自定义类型
	aboutType()
	
	// 2.结构体 - 基本实例化
	aboutMyStruct()

	// 3.指针类型结构体
	aboutPtrStruct()

	// 4. 结构体初始化
	aboutStructInit()

	// 5. struct 内存布局
	aboutStructMemoryLayout()

	// 6.面试题
	aboutStructInterview()

	// 7.构造函数
	aboutStructTest7()

	// 8. 方法和接受者
	aboutStructTest8()

	// 9. 指针类型接收者
	aboutStructTest9()

	// 10.值类型接收者
	aboutStructTest10()

	// 11.0 结构体匿名字段
	aboutStructTest11()

	// 12.结构体与JSON序列化
	aboutStructTest12()

	// 13. 结构体标签
	aboutStructTest13()
}

// 1.类型别名和自定义类型
func aboutType() {
	// 1.1自定义类型
	// 通过Type关键字的定义，MyInt就是一种新的类型，它具有int的特性。
	// 类型定义
	// type MyInt int

	// 1.2类型别名

	// 类型别名
	// type MyInt = int

	// 1.3 类型定义与类型别名的区别
	var a NewInt
	var b MyInt
	fmt.Printf("type of a：%T\n", a) // type of a：main.NewInt
	fmt.Printf("type of b：%T\n", b) // type of b：int

	// 结论：结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。
	// MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
}

// 2.结构体

type person struct {
	name string
	city string
	// name, city string // 同类型的字段也可以写在一行
	age int8
}

func aboutMyStruct() {
	// 2.1结构体实例化（初始化）
	// 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段
	var p1 person
	fmt.Printf("p1 = %#v\n", p1)
	p1.name = "Alex"
	p1.city = "北京"
	p1.age = 20
	fmt.Printf("p1=%v\n", p1)   // p1={Alex 北京 20}
	fmt.Printf("p1=%#v\n", p1)  // p1=main.person{name:"Alex", city:"北京", age:20}。 %#v：输出go语言语法结构的值

	// 2.2 匿名结构体
	var user struct{name string; age int}
	user.name = "Tom"
	user.age = 19
	fmt.Printf("user = %#v\n", user)
}

// 3.结构体指针
func aboutPtrStruct() {

	// 3.1 使用 new 实例化结构体，获取结构体地址
	// 我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：
	// p2 是结构体指针
	var p2 = new(person)
	fmt.Printf("type of p2 = %T\n", p2) // type of p2 = *main.person
	fmt.Printf("p2 = %#v\n", p2)		// p2 = &main.person{name:"", city:"", age:0}

	// 注意：Go语言中支持对结构体指针直接使用.来访问结构体的成员。（区别于C语言）
	p2.name = "Pony"
	p2.age = 45
	p2.city = "深圳"
	fmt.Printf("p2 = %#v\n", p2) 		// p2 = &main.person{name:"Pony", city:"深圳", age:45}

	// 3.2 取结构体的地址实例化
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	var p3 = &person{}
	fmt.Printf("type of p3 = %T\n", p3) // type of p3 = *main.person
	fmt.Printf("p3 = %#v\n", p3)		// p3 = &main.person{name:"", city:"", age:0}

	// p3.name = "Jack" 其实在底层是(*p3).name = "Jack"，这是Go语言帮我们实现的语法糖。
	p3.name = "Jack"
	p3.age = 19
	p3.city = "珠海"
	fmt.Printf("p3 = %#v\n", p3) 		// p3 = &main.person{name:"Jack", city:"珠海", age:19}
}

// 4. 结构体初始化
func aboutStructInit() {
	// 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
	// 4.1 使用键值对初始化
	p4 := person{
		name: "Rose",
		city: "厦门",
		age: 22,
	}
	fmt.Printf("4.1 == p4 = %#v\n", p4) // 4.1 == p4 = main.person{name:"Rose", city:"厦门", age:22}

	// 4.2 对结构体指针进行键值对初始化
	p5 := &person{
		name: "Fill",
		city: "首尔",
		age: 66,
	}
	fmt.Printf("4.2 == p5 = %#v\n", p5) // 4.2 == p5 = &main.person{name:"Fill", city:"首尔", age:66}

	// 初始化部分字段
	p6 := &person{
		city: "南京",
	}
	fmt.Printf("4.2 == p6 = %#v\n", p6) // 4.2 == p6 = &main.person{name:"", city:"南京", age:0}


	// 4.3 使用值的列表初始化
	/*
	该方式使用注事事项：
		1.必须初始化结构体的所有字段。
    	2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
    	3.该方式不能和键值初始化方式混用。
	*/
	p7 := &person{
		"Zhang",
		"咸阳",
		12,
	}
	fmt.Printf("4.3 == p7 = %#v\n", p7)
}


// 5. struct 内存布局
type testStruct struct {
	a int8
	b int8
	c int8
	d int8
}
func aboutStructMemoryLayout() {
	n := testStruct{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	/*
		n.a 0xc0000c4118
		n.b 0xc0000c4119
		n.c 0xc0000c411a
		n.d 0xc0000c411b
	*/
}

// 6.面试题
type student struct {
    name string
    age  int
}

func aboutStructInterview() {
	m := make(map[string]*student)
    stus := []student{
        {name: "pprof.cn", age: 18},
        {name: "测试", age: 23},
        {name: "博客", age: 28},
    }

    for _, stu := range stus {
        m[stu.name] = &stu
    }
    for k, v := range m {
        fmt.Println(k, "=>", v.name)
    }
}

// 7.构造函数
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newPerson(name, city string, age int8) *person {
    return &person{
        name: name,
        city: city,
        age:  age,
    }
}
func aboutStructTest7() {
	p9 := newPerson("Tony", "北京", 19)
	fmt.Printf("%#v\n", p9)
}

// 8. 方法和接受者
/*
	Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。
	方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。

	方法的定义格式如下：
    func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
        函数体
    }

 	1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
    2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
    3.方法名、参数列表、返回参数：具体格式与函数定义相同。
*/

// 构造函数
func NewPerson(name string, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age: age,
	}
}

// 做梦
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func aboutStructTest8() {
	p1 := newPerson("Alex", "北京", 22)
	p1.Dream()
}

// 9. 指针类型接收者
/*
什么时候应该使用指针类型接收者？？？
    1.需要修改接收者中的值
    2.接收者是拷贝代价比较大的大对象
    3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/
func (p *person)SetAge(newAge int8) {
	p.age = newAge
}

func aboutStructTest9() {
	p1 := newPerson("Alex", "北京", 22)
	fmt.Printf("p1.age = %v\n", p1.age) // p1.age = 22
	p1.SetAge(30)
	fmt.Printf("p1.age = %v\n", p1.age) // p1.age = 30
}

// 10.值类型接收者
func (p person)SetAge2(newAge int8) {
	p.age = newAge
}
func aboutStructTest10() {
	p1 := newPerson("Alex", "北京", 22)
	fmt.Printf("p1.age = %v\n", p1.age) // p1.age = 22
	p1.SetAge2(30)
	fmt.Printf("p1.age = %v\n", p1.age) // p1.age = 22
}

// 11.0 结构体匿名字段
// 匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
type stu struct {
	string
	int8
}
func aboutStructTest11() {
	p1 := stu{
		"abc",
		10,
	}
	fmt.Printf("p1 = %#v\n", p1)
}

// 结构体嵌套
// 结构体嵌套匿名结构体
// 结构体的集成：通过嵌套匿名结构体实现继承
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

// 12.结构体与JSON序列化
type Student struct {
	ID int
	Gender string
	Name string
}

type Class struct {
	Title string
	Student []*Student // 指针数组
}

func aboutStructTest12() {
	fmt.Println("===12.结构体与JSON序列化===")

	c := &Class{
		Title: "101",
		Student: make([]*Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			ID: i,
			Gender: "男",
			Name: fmt.Sprintf("stu%02d", i),
		}
		c.Student = append(c.Student, stu)
	}

	// JSON序列化：结构体-->JSON格式的字符串
	data, error := json.Marshal(c)
	if error != nil {
		fmt.Println("JSON序列化失败。。。")
		return
	}
	fmt.Printf("data: %s\n", data)

	// JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Student":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}` 
	c1 := &Class{}
	err := json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("JSON反序列化。。。")
		return
	}
	fmt.Printf("c1 = %#v\n", c1)
}

// 13. 结构体标签
// Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。
/*

Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
	`key1:"value1" key2:"value2"`
结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。
*/
type Animal struct {
	ID int			`json:"id"` // 通过指定tag实现json序列化该字段时的key
	Gender string  // json序列化是默认使用字段名作为key
	name string	   // 私有不能被json包访问
}

func aboutStructTest13() {
	s1 := Animal{
		ID: 1,
		Gender: "雄性",
		name: "FF",
	}

	data, error := json.Marshal(s1)
	if error != nil {
		fmt.Println("JSON序列化失败")
		return
	}
	fmt.Printf("data = %s\n", data) // data = {"id":1,"Gender":"雄性"}
}