package main

import (
	"fmt"
)

func aboutMethod() {

	// receiver T 和 *T 的差别
	fmt.Println("方法接收者：receiver T 和 *T 的差别")
	methodTest1()

	// 函数与方法的区别
	fmt.Println("函数与方法的区别")
	fmt.Println("===函数===")
	structTestValue()
	fmt.Println("===方法===")
	structTestFunc()


    // 表达式
    fmt.Println("===表达式===")
    aboutExpression()
	aboutExpression2()

    // 依据方法集转换 method expression，注意 receiver 类型的差异。
    aboutExpression3()
}

/*===================================================*/
//                  方法定义                           /
/*===================================================*/

/*
Golang 方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)。

• 只能为当前包内命名类型定义方法。
• 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。 
• 不支持方法重载，receiver 只是参数签名的组成部分。
• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。

定义：
    func (recevier type) methodName(参数列表)(返回值列表){}
    注：参数和返回值可以省略
*/

// receiver T 和 *T 的差别
type Data struct {
    x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
    fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() { // func PointerTest(self *Data);
    fmt.Printf("Pointer: %p\n", self)
}

func methodTest1() {
    d := Data{}
    p := &d
    fmt.Printf("Data: %p\n", p) // Data: 0xc000018320

    d.ValueTest()   // ValueTest(d)      0xc000018328
    d.PointerTest() // PointerTest(&d)   0xc000018320

    p.ValueTest()   // ValueTest(*p)     0xc000018330
    p.PointerTest() // PointerTest(p)    0xc000018320

	/*
	结论：
		当方法接收者是指针时，即使用值类型调用那么函数内部也是对指针的操作
		当方法接收者不是一个指针时，该方法操作对应接收者的值的副本(意思就是即使你使用了指针调用函数，但是函数的接收者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。
	*/ 
}



/*===================================================*/
//                  普通函数与方法的区别                 /
/*===================================================*/

/*
1.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递。反之亦然。
2.对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法。反过来同样也可以。
*/

//1.普通函数
//接收值类型参数的函数
func valueIntTest(a int) int {
    return a + 10
}

//接收指针类型参数的函数
func pointerIntTest(a *int) int {
    return *a + 10
}

func structTestValue() {
    a := 2
    fmt.Println("valueIntTest:", valueIntTest(a))
    //函数的参数为值类型，则不能直接将指针作为参数传递
    //fmt.Println("valueIntTest:", valueIntTest(&a))
    //compile error: cannot use &a (type *int) as type int in function argument

    b := 5
    fmt.Println("pointerIntTest:", pointerIntTest(&b))
    //同样，当函数的参数为指针类型时，也不能直接将值类型作为参数传递
    //fmt.Println("pointerIntTest:", pointerIntTest(b))
    //compile error:cannot use b (type int) as type *int in function argument
}

//2.方法
type PersonD struct {
    id   int
    name string
}

//接收者为值类型
func (p PersonD) valueShowName() {
    fmt.Println(p.name)
}

//接收者为指针类型
func (p *PersonD) pointShowName() {
    fmt.Println(p.name)
}

func structTestFunc() {
    //值类型调用方法
    personValue := PersonD{101, "hello world"}
    personValue.valueShowName()
    personValue.pointShowName()

    //指针类型调用方法
    personPointer := &PersonD{102, "hello golang"}
    personPointer.valueShowName()
    personPointer.pointShowName()

    //与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
}


/*===================================================*/
//                     匿名字段                        /
/*===================================================*/


/*===================================================*/
//                      方法集                         /
/*===================================================*/
/*
Golang方法集 ：每个类型都有与之关联的方法集，这会影响到接口实现规则。

Go 语言中内部类型方法集提升的规则：
    • 类型 T 方法集包含全部 receiver T 方法。
    • 类型 *T 方法集包含全部 receiver T + *T 方法。
    • 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。 
    • 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。 
    • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。

用实例 value 和 pointer 调用方法 (含匿名字段) 不受方法集约束，编译器总是查找全部方法，并自动转换 receiver 实参。    
*/


/*===================================================*/
//                      表达式                         /
/*===================================================*/
/*
Golang 表达式 ：根据调用者不同，方法分为两种表现形式:
    1、instance.method(args...)，称为 method value
    2、<type>.func(instance, args...)，称为 method expression
区别在于 method value 绑定实例，而 method expression 则须显示传参
*/

type User struct {
    id   int
    name string
}

func (self *User) Test() {
    fmt.Printf("%p, %v\n", self, self)
}

// 表达式
func aboutExpression() {
    fmt.Println("===aboutExpression===")

    u := User{1, "Tom"}
    u.Test()                             // 0xc0000b0138, &{1 Tom}

    mValue := u.Test 
    u.id = 2
    mValue() // 隐式传递 receiver            0xc0000b0138, &{2 Tom}

    mExpression := (*User).Test
    mExpression(&u) // 显示传递 receiver     0xc0000b0138, &{2 Tom}
}

// 值类型

func (self User) Test2() {
    fmt.Printf("%p, %v\n", &self, self)
}

func aboutExpression2() {
    fmt.Println("===aboutExpression2===")
    u := User{1, "Tom"}
    mValue := u.Test2 // 需要注意，method value 会复制 receiver，因为不是指针类型，不受后续修改影响。
    // 在汇编层面，method value 和闭包的实现方式相同，实际返回 FuncVal 类型对象。
    // FuncVal { method_address, receiver_copy }

    u.id, u.name = 2, "Jack"
    u.Test2()

    mValue()
}


// 依据方法集转换 method expression，注意 receiver 类型的差异。

func (self *User) TestPointer() {
    fmt.Printf("TestPointer: %p, %v\n", self, self)
}

func (self User) TestValue() {
    fmt.Printf("TestValue: %p, %v\n", &self, self)
}

func aboutExpression3() {
    fmt.Println("===aboutExpression3===")

    u := User{1, "Tom"}
    fmt.Printf("User: %p, %v\n", &u, u)  // User: 0xc0001281f8, {1 Tom}

    mv := User.TestValue
    mv(u)                                // TestValue: 0xc000128228, {1 Tom}

    mp := (*User).TestPointer
    mp(&u)                               // TestPointer: 0xc0001281f8, &{1 Tom}

    // *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
    mp2 := (*User).TestValue 
    mp2(&u)                              // TestValue: 0xc000128270, {1 Tom}
}


// 将方法 "还原" 成函数，就容易理解下面的代码了。
type MyData struct{}
func (MyData) TestMyValue() {}
func (*MyData) TestMyPointer() {}

func aboutExpression4() {
    var p *MyData = nil
    p.TestMyPointer()

    (*MyData)(nil).TestMyPointer() // method value
    (*MyData).TestMyPointer(nil)   // method expression
    
    // p.TestMyValue()             // invalid memory address or nil pointer dereference
    // (MyData)(nil).TestMyValue() // cannot convert nil to type Data
    // Data.TestValue(nil)         // cannot use nil as type Data in function argument
}