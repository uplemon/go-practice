package main

import "fmt"

func main() {
    varTest()
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


