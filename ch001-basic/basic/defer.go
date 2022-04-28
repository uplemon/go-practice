package basic

import (
    "fmt"
    "io"
    "os"
)


func DeferDemo() {
    filename := "/tmp/tmp.txt"
    ReadFile(filename)
    MultiDeferDemo()
}

/**
 * defer关键字用于修饰一个函数或者方法，使得该函数或方法在返回前才会执行，也就是说会延迟但又保证一定执行
 * defer语句常被用于成对的操作，如文件的打开和关闭，加锁和释放锁，连接的建立和断开等
 * 不管多么复杂的操作，都可以保证资源被正确的释放
 */
func ReadFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("file is not exists")
        }
        return nil, err
    }
    defer f.Close()
    contentBytes, err := io.ReadAll(f)
    if err == nil {
        fmt.Println(string(contentBytes))
    }
    return contentBytes, err
}

/**
 * 多个defer的定义与执行类似于栈的操作：先进后出，最先定义的最后执行
 * 输出结果如下:
 * main_func[x=>0]
 * defer_func_3[x=>1]
 * defer_func_2[x=>2]
 * defer_func_1[x=>3]
*/
func MultiDeferDemo()  {
    x := 0
    defer func() {
        x++
        fmt.Printf("defer_func_1[x=>%d]\n", x)
    }()
    defer func() {
        x++
        fmt.Printf("defer_func_2[x=>%d]\n", x)
    }()
    defer func() {
        x++
        fmt.Printf("defer_func_3[x=>%d]\n", x)
    }()
    fmt.Printf("main_func[x=>%d]\n", x)
}