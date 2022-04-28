package basic

import "fmt"

/**
 * 接口是和调用方的一种约定，不用和具体的实现细节绑定在一起
 * 接口要做的是定义好约定，告诉调用方自己可以做什么，但无需关心其内部实现
 */
type Mobile interface {
    call() string
}

type Huawei struct {}
type Xiaomi struct {}
type Apple struct {}

/**
 * Go中的接口实现与Java中不一样，Java是声明的，必须使用关键字implements
 * Go用的是Duck-Like模式，如果类型中定义的方法与接口中的一致，就代表该类型实现了该接口
 * 如果一个接口有多个方法，那么需要实现接口的每个方法才算是实现了这个接口
 */
func (m Huawei) call() string {
    return "Huawei"
}

func (m Xiaomi) call() string {
    return "Xiaomi"
}

func (m *Apple) call() string {
    return "Apple"
}

/**
 * 以值类型接收者实现接口的时候，不管是类型本身，还是该类型的指针类型，都实现了该接口
 * 以指针类型接收者实现接口的时候，只有对应的指针类型才被认为实现了该接口
 */
func InterfaceDemo() {
    var h1 Mobile = Huawei{}
    fmt.Println(h1.call())
    var h2 Mobile = new(Huawei)
    fmt.Println(h2.call())

    var x1 Mobile = Xiaomi{}
    fmt.Println(x1.call())
    var x2 Mobile = new(Xiaomi)
    fmt.Println(x2.call())

    // 只有Apple的指针类型实现了该接口
    // var a1 Mobile = Apple{}
    // fmt.Println(a1.call())
    var a2 Mobile = new(Apple)
    fmt.Println(a2.call())
}
