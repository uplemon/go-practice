package quickstart

import "fmt"

// 在Go语言中，方法和函数是两个概念，但又非常相似，不同点在于方法必须要有一个接收者
// 这个接收者是一个类型，这样方法就和这个类型绑定在一起，称为这个类型的方法

type Age uint

// 定义方法会在关键字func和方法名之间加一个接收者，接收者使用小括号包围
// 接收者的定义和普通变量、函数参数等一样，前面是变量名，后面是接收者类型
func (age Age) String() {
    fmt.Println("age:", age)
}

// 值接收者
func (age Age) Modify1() {
    age = Age(30)
}

// 指针接收者
func (age *Age) Modify2() {
    *age = Age(30)
}

/**
 * 方法的接收者可以是值类型，也可以是指针类型
 * 如果接收者是指针类型，我们对指针的修改是有效的，如果不是指针类型，修改就没有效果
 */
func MethodDemo() {
    age := Age(20)
    age.String()
    age.Modify1()
    age.String()
    age.Modify2()
    age.String()
}
