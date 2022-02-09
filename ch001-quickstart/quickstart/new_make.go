package quickstart

import (
    "fmt"
    "reflect"
)

func NewMakeDemo() {
    //varTest()
    newTest()
    //varInitTest()
    makeTest()
}

/**
 * 变量的声明，初始化会涉及到内存的分配
 * 对于值类型，即使只声明变量，没有对其进行初始化，该变量也会有分配好的内存
 * 对于指针类型，声明后默认是零值nil，该变量没有指向的内存空间，如果进行赋值操作就会引发nil指针错误
 * 总结：如果要对一个变量赋值，这个变量必须有对应分配好的内存，这样才能对这块内存操作完成赋值目的
 */
func varTest()  {
    var str string
    str = "hello str"
    fmt.Println(str)

    var strP *string
    // panic: runtime error: invalid memory address or nil pointer dereference
    *strP = "hello strP"
    fmt.Println(*strP)

    var strP1 *string = &str
    *strP1 = "hello strP1"
    fmt.Println(*strP1)
}

/**
 * 声明指针变量默认是没有分配内存的，可以通过内置函数new给它分配一块内存
 * new的作用就是根据传入的类型申请一块内存，然后返回指向这块内存的指针，指针指向的数据就是该类型的零值
 * 内置new函数定义：func new(Type) *Type
 */
func newTest() {
    var strP *string
    strP = new(string)
    // 打印空字符串，也就是string的零值
    fmt.Println(*strP)
    // new已为指针变量分配了内存，可以直接赋值
    *strP = "hello strP"
    fmt.Println(*strP)
}

/**
 * 定义一个结构体类型
 */
type Student struct {
    name string
    age int32
}
/**
 * 变量初始化（不初始化的变量的值为该变量类型的零值）
 * 当声明一个类型的变量时还对这个变量进行了赋值，这个修改变量值的过程称为变量的初始化
 */
func varInitTest() {
    // 字面量初始化，基础类型和复合类型都可以通过这种方式进行初始化
    p1 := Student{name: "student1", age: 19}
    // 指针变量初始化
    p2 := newStudent("student2", 20)
    // 值变量
    fmt.Printf("p1 type: %T\n", p1)
    fmt.Println("p1 type:", reflect.TypeOf(p1))
    // 指针变量
    fmt.Printf("p2 type: %T\n", p2)
    fmt.Println("p2 type:", reflect.TypeOf(p2))
}
// 通过封装函数初始化变量（工厂函数）
func newStudent(name string, age int32) *Student {
    s := new(Student)
    s.name = name
    s.age = age
    return s
}

/**
 * make返回引用类型
 * make函数只用于slice、map、chan这三种内置类型的创建和初始化
 */
func makeTest() {
    makeSliceTest()
    makeMapTest()
}

func makeSliceTest() {
    slice := make([]string, 3, 3)
    slice[0] = "a"
    slice[1] = "b"
    slice[2] = "c"
    fmt.Println(slice)
}

func makeMapTest() {
    // mapVar := map[string]string{"beijing":"北京", "shanghai":"上海"}
    var m map[string]string
    m = make(map[string]string)
    // m := make(map[string]string)
    m["beijing"] = "北京"
    m["shanghai"] = "上海"
    m["guangzhou"] = "广州"
    m["shenzhen"] = "深圳"
    m["hangzhou"] = "杭州"
    fmt.Println("m type:", reflect.TypeOf(m))
    fmt.Println(m)
}

// 函数new和make的区别？
// 1.new函数只用于分配内存，并且把内存清零，也就是返回一个指向对应类型零值的指针。new函数一般用于需要显式返回指针的情况。
// 2.make函数只用于slice、chan和map这三种内置类型的创建和初始化。
