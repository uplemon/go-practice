package basic

import (
    "fmt"
    "math/rand"
)

// if...else if...else条件语句
func IfDemo() {
    min, max := 10, 60
    // 模拟随机生成[min, max)
    num := rand.Intn(max - min) + min
    fmt.Println(num)
    if num >=10 && num < 20 {
        fmt.Println("10<=num<20")
    } else if num >= 20 && num < 30 {
        fmt.Println("20<=num<30")
    } else if num >= 30 && num < 40 {
        fmt.Println("30<=num<40")
    } else if num >= 40 && num < 50 {
        fmt.Println("40<=num<50")
    } else {
        fmt.Println("50<=num<60")
    }
}

// switch选择语句
func SwitchDemo()  {
    min, max := 10, 60
    switch n := rand.Intn(max - min) + min; {
    case n >=10 && n < 20 :
        fmt.Println("10<=n<20")
    case n >= 20 && n < 30:
        fmt.Println("20<=n<30")
    case n >= 30 && n < 40:
        fmt.Println("30<=n<40")
    case n >= 40 && n < 50:
        fmt.Println("40<=n<50")
    default:
        fmt.Println("50<=n<60")
    }

    // fallthrough可以执行下一个紧跟的case
    // case后的值要和i结果类型相同，这里的i是int类型，case后的值就只能使用int类型...
    // ...如果使用其他类型，会提示类型不匹配，无法编译通过
    switch i:=1;i {
    case 1:
        fallthrough
    default:
        fmt.Println(1)
    }
}

func ForDemo()  {
    // 计算min~max之间所有数字之和
    min, max := 0, 100
    sum := 0
    for i:=min;i<=max;i++ {
        sum += i
    }
    fmt.Println("sum:", sum)

    // Go语言中没有while循环，但可以用for达到while的效果
    min, max, sum, i := 0, 100, 0, min
    for i<=max {
        sum += i
        i++
    }
    fmt.Println("sum:", sum)
}
