package basic

import (
	"errors"
	"fmt"
)

func FuncDemo() {
	fmt.Println("FuncDemo1:", FuncDemo1(1, 2))

	sum, err := FuncDemo2(2, 3)
	if err == nil {
		fmt.Println("FuncDemo2:", sum)
	}

	sum1, err1 := FuncDemo3(3, 4)
	if err1 == nil {
		fmt.Println("FuncDemo3:", sum1)
	}

	sum2 := FuncDemo4(1, 2, 3, 4, 5)
	fmt.Println("FuncDemo4:", sum2)

	FuncDemo5(5, 6)

	FuncDemo6()
}

/**
 * 标准函数构成：
 * 1. 关键字func
 * 2. 函数名字funcName
 * 3. 函数参数
 * 4. 函数返回值
 * 5. 函数体
 */
func FuncDemo1(a, b int) int {
	return a + b
}

// 多值返回函数，Go语言的函数可以返回多个值
// 在Go语言标准库中有很多这样的函数，第一个值返回函数的结果，第二个值返回函数出错的信息，这就是多值返回
func FuncDemo2(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("参数必须是正数")
	}
	return a + b, nil
}

// 命名返回值
// 可以为每个返回值起一个名字，这个名字可以像参数一样在函数体内使用
func FuncDemo3(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("参数必须是正数")
	}
	sum = a + b
	err = nil
	return
}

// 可变参数，即函数的参数数量是可变的
// 同一个函数，可以不传参数，也可以传递一个参数，也可以传递多个参数，这种函数就是具有可变参数的函数
// 例如最常见的fmt.Println函数：func Println(a ...interface{}) (n int, err error)
// 定义可变参数，只要在参数类型前加三个点...即可
func FuncDemo4(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

// 匿名函数
func FuncDemo5(a, b int) {
	// sum对应的值就是一个匿名函数，这里的sum只是一个函数类型的变量，并不是函数的名字
	sum := func(x, y int) int {
		return x + y
	}
	fmt.Println("FuncDemo5:", sum(a, b))
}

// 闭包
// 有了匿名函数，就可以在函数中再定义函数（函数嵌套），定义的这个匿名函数，也可以称为内部函数
// 在函数内定义的内部函数，可以使用外部函数的变量等，这种方式称为闭包
func FuncDemo6() {
	cl := closure()
	// 每调用一次cl()，i的值就会加1
	fmt.Println("FuncDemo6:", cl())
	fmt.Println("FuncDemo6:", cl())
	fmt.Println("FuncDemo6:", cl())
}

// 得益于闭包函数闭包的能力，自定义closure函数，可以返回一个匿名函数并且持有外部函数closure的变量i
func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
