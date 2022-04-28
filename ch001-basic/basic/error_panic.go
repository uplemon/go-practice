package basic

import (
    "errors"
    "fmt"
    "strconv"
)

func ErrorDemo() {
    errorStrToIntTest()
    errorSumTest()
    commonErrorSumTest()
}

func PanicDemo() {
    connDB("", "root", "123456")
}

/**
 * 在Go语言中错误是可以预期的，并且不是非常严重，不会影响程序的运行
 * 对于这类问题，可以用返回错误给调用者的方式，让调用者自己决定如何处理
 * error接口：在Go语言中错误是通过内置的error接口表示的，该接口只有一个Error方法用来返回具体的错误信息
 * type error interface {
 *     Error() string
 * }
 */

/**
 * 来看一个字符串转整数的例子，因为字符串"a"无法转为整数，所以这段程序会打印如下错误信息
 * strconv.Atoi: parsing "a": invalid syntax
 * 这个错误信息就是通过接口error返回的，strconv.Atoi函数的定义:func Atoi(s string) (int, error)
 * error接口用于当方法或者函数执行遇到错误时进行返回，而且是第二个返回值
 * 通过这种方式，可以让调用者自己根据错误信息决定如何进行下一步处理
 */
func errorStrToIntTest() {
    i, err := strconv.Atoi("a")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(i)
    }
}

/**
 * error工厂函数: errors.New(string)
 * 自定义函数也可以返回错误信息给调用者
 */
func errorSum(a, b int) (int, error) {
    if a<0 || b<0 {
        return 0, errors.New("a或者b不能为负数")
    } else {
        return a+b, nil
    }
}
func errorSumTest() {
    sum, err := errorSum(-1, 2)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(sum)
    }
}

/**
 * 自定义error
 * 上面采用工厂函数返回错误信息的方式只能传递一个字符串，如果想要传递更多的信息(比如错误码)，这时就需要自定义error
 * 自定义error其实就是先自定义一个新类型，比如结构体，然后让这个类型实现error接口
 */
type commonError struct {
    errorCode int // 错误码
    errorMsg string // 错误信息
}
// 实现error接口
func (ce *commonError) Error() string {
    return ce.errorMsg
}
// 返回自定义error(返回更多信息)
func commonErrorSum(a, b int) (int, error) {
    if a<0 || b<0 {
        return 0, &commonError{
            errorCode: 1,
            errorMsg: "a或者b不能为负数",
        }
    } else {
        return a+b, nil
    }
}
/**
 * 有了自定义error，并且可以包含更多的错误信息后，就可以使用这些信息了
 * 需要先把返回的error接口转换为自定义的错误类型(使用类型断言)
 */
func commonErrorSumTest() {
    sum, err := commonErrorSum(-1, 2)
    if cm, ok := err.(*commonError); ok {
        fmt.Printf("errorCode:%d, errorMsg:%s\n", cm.errorCode, cm.errorMsg)
    } else {
        fmt.Println(sum)
    }
}

/**
 * Panic异常
 * Go语言是一门静态的强类型语言，很多问题都尽可能在编译时捕获，但是有一些只能在运行时检查
 * 比如数组越界访问，不相同的类型强制转换，这类运行时的问题会引起panic异常
 * 除了运行时可以产生panic外，我们自己也可以抛出panic异常
 * panic是Go语言内置的函数，可以接受interface{}类型的参数，也就是任意类型的值都可以传递给panic函数
 * func panic(v interface{})
 * 提示：interface{}是空接口的意思，在Go语言中代表任意类型
 * panic异常是一种非常严重的情况，会让程序中断运行，使程序崩溃
 * 如果是不影响程序运行的错误，不要使用panic，使用普通错误error即可
 *
 * Recover捕获Panic异常
 * 通常情况下我们不对panic异常做任何处理，因为既然它是影响程序运行的异常，就让它直接崩溃即可
 * 但也有一些例外，比如程序崩溃前做一些资源释放的处理，这就需要从panic异常中恢复才能完成处理
 * 可以通过内置的recover函数恢复panic异常
 * 因为在程序panic异常崩溃的时，只有被defer修饰的函数才能被执行，所以recover函数要结合defer关键字使用才能生效
 * [defer关键字 + 匿名函数 + recover函数]从panic异常中恢复
 */
func connDB(host, username, password string) {
    defer func() {
        if p := recover(); p != nil {
            fmt.Println(p)
            //...
        }
    }()
    if host == "" || username == "" || password == "" {
        panic("(host|username|password)不能为空")
    }
    //...
}
