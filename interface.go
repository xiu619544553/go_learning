package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
1、接口类型：
	在Go语言中接口（interface）是一种类型，一种抽象的类型。为了保护你的Go语言职业生涯，请牢记接口（interface）是一种类型。
	interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

2、为什么使用接口
	接口区别于我们之前所有的具体类型，接口是一种抽象的类型。当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么。
	比如：动物的叫、吃等行为，通过接口来定义

3、接口的定义：Go语言提倡面向接口编程。


接口命名习惯以 er 结尾。

每个接口由数个方法组成，接口的定义格式如下：
    type 接口类型名 interface{
        方法名1( 参数列表1 ) 返回值列表1
        方法名2( 参数列表2 ) 返回值列表2
        …
    }
1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
*/

// 举个例子
type writer interface {
	Write([]byte) error
}

func aboutInterface() {
	fmt.Println("go interface")

	// 值接收者
	valueReceiveImpInterface()

	// 指针接收者实现接口
	ptrImpInterface()

	// 空接口作为map的值
	emptyInterface()

	// 类型断言
	styleAsset()
}

// 示例
type Sayer interface {
	say()
}

type Dog struct {}
type Cat struct {}

// 因为Sayer接口里只有一个say方法，所以我们只需要给dog和cat 分别实现say方法就可以实现Sayer接口了。
func (d Dog) say() {
	fmt.Println("汪汪~")
}

func (c Cat) say() {
	fmt.Println("喵喵~")
}


// 接口类型变量能够存储所有实现了该接口的实例。
func aboutImpInterface() {
	var x Sayer // 声明一个Sayer类型的变量x
    a := Cat{}  // 实例化一个cat
    b := Dog{}  // 实例化一个dog
    x = a       // 可以把cat实例直接赋值给x
    x.say()     // 喵喵喵
    x = b       // 可以把dog实例直接赋值给x
    x.say()     // 汪汪汪
}


// 值接收者和指针接收者实现接口的区别

// 值接收者
func valueReceiveImpInterface() {
	var x Mover
	var wangcai = Dog{}
	x = wangcai
	var fugui = &Dog{}
	x = fugui
	x.move()

	/*
	从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。
	*/
}

type Mover interface {
	move()
}

func (d Dog) move() {
	fmt.Println("狗会动")
}

// 指针接收者实现接口
func ptrImpInterface() {
	// var x Mover
    // var wangcai = Dog{} // 旺财是dog类型
    // x = wangcai         // x不可以接收dog类型
    // var fugui = &Dog{}  // 富贵是*dog类型
    // x = fugui           // x可以接收*dog类型

	// 此时实现 Move2r 接口的是 *Dog 类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。
}

type Move2r interface {
	move()
}

func (d *Dog) move2() {
	fmt.Println("狗会动")
}


// 类型与接口的关系
// 一个类型实现多个接口
// 多个类型实现同一接口：比如 Dog、Cat 都实现了 Sayer 接口\
// 接口嵌套
type animaler interface {
	Sayer  // 接口 Sayer
	Mover  // 接口 Mover
}


// 空接口的定义
// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量。


// 空接口的应用
// 空接口作为函数的参数
// 使用空接口实现可以接收任意类型的函数参数

// 空接口作为函数参数
func show(a interface{}) {
    fmt.Printf("type:%T value:%v\n", a, a)
}


// 空接口作为map的值
// 使用空接口实现可以保存任意值的字典。
func emptyInterface() {
	fmt.Println("===空接口作为map的值===")
    var studentInfo = make(map[string]interface{})
    studentInfo["name"] = "李白"
    studentInfo["age"] = 18
    studentInfo["married"] = false
    fmt.Println(studentInfo)
}


// 类型断言
// 接口值
// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。
func styleAsset() {
	fmt.Println("===类型断言===")
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = nil
	fmt.Println(w)	

	// 想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：x.(T)
	// x：表示类型为interface{}的变量
    // T：表示断言x可能是的类型。
	// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
	var x interface{}
	x = "www.baidu.com"
	v, ok := x.(string)
	if ok { // 断言类型正确
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	var a = []byte("abc")
	fmt.Println(a)
	b := bytes.ToUpper(a)
	fmt.Println(b)
}