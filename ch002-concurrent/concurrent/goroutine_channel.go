package concurrent

import (
    "fmt"
    "math/rand"
    "time"
)

/**
 * Go语言中没有线程的概念，只有协程，也称为goroutine
 * 相比线程来说，协程更加轻量，一个程序可以随意启动成千上万个goroutine
 * goroutine被Go runtime所调用，这一点和线程不一样
 * Go语言的并发是由Go自己所调度的，自己决定同时执行多少个goroutine，什么时候执行哪几个
 * goroutine的调度对于开发者完全透明，开发者只需要在编码时告诉Go语言要启动几个goroutine
 * 启动一个goroutine非常简单，使用go关键字，
 */
func GoroutineDemo()  {
    goroutineDemo1()
    goroutineDemo2()
    goroutineDemo3()
    selectDemo()
}

/**
 * 程序是并发的，go关键字启动的goroutine并不阻塞main goroutine的执行
 * time.Sleep()表示等待，不然main goroutine执行完毕程序就会退出，就看不到新启动的goroutine执行结果
 */
func goroutineDemo1() {
    go func(name string) {
        fmt.Println(name)
    }("goroutine-1")

    go fmt.Println("goroutine-2")

    go func() {
        fmt.Println("goroutine-3")
    }()
    fmt.Println("main goroutine")
    time.Sleep(time.Second)
}

/**
 * 如果启动了多个goroutine，它们之间该如何通信呢？这就是Go语言提供的channel(通道)要解决的问题
 * 声明一个channel，使用内置的make函数即可: ch := make(chan string)
 * 其中chan是一个关键字，表示channel类型。后面的string表示channel里的数据是string类型
 * 通过channel的声明可以看到，chan是一个集合类型
 * 一个chan的操作只有两种:发送和接收
 * 1. 接收：获取chan中的值，操作符为<-chan
 * 2. 发送：向chan发送值，即把值放到chan中，操作符为chan<-
 * 这里注意发送和接收的操作符都是<-，接收的<-操作符在chan的左侧，发送的<-操作符在chan的右侧
 * channel有点像在两个goroutine之间架设的管道，一个可以往这个管道里发送数据，另一个可以从这个管道取数据
 */
func goroutineDemo2() {
    ch := make(chan string)
    // 启动新的goroutine向channel中发送值
    go func() {
        fmt.Println("goroutine-1")
        ch <- "[goroutine-1]执行完成"
    }()
    fmt.Println("main goroutine")
    // 在main goroutine中从channel中接收值，如果channel没有值则阻塞等到channel中有值可以接收为止
    v := <-ch
    fmt.Println("接收到channel中的值为:", v)
}

/**
 * 无缓冲channel/有缓冲channel
 * 使用make创建的chan就是一个无缓冲channel(容量为0)，它不能存储任何数据，只起到传输数据作用
 * 无缓冲channel的发送和接收操作是同时进行的，它也可以称为同步channel
 * 有缓冲channel类似一个可阻塞的队列，内部的元素先进先出。通过make函数的第二个参数可以指定channel容量的大小，进而创建一个有缓冲channel
 * ch := make(chan string, 5)
 * 如上所示，创建了一个容量为5的队列，内部元素类型是string，也就是说这个channel内部最多可以存放5个类型为string的元素
 * 一个有缓冲的channel具备以下特点：
 * 1. 有缓冲channel的内部有一个缓冲队列
 * 2. 发送操作是向队列的尾部插入元素，如果队列已满则阻塞等待，直到另一个goroutine执行接收操作释放队列的空间
 * 3. 接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空则阻塞等待，直到另一个goroutine执行发送操作插入新的元素
 *
 * 关闭channel
 * channel可以使用内置函数close关闭，如close(ch)
 * 如果一个channel被关闭了，就不能向里面发送数据了，如果发送的话会引起panic异常
 * 但还可以接收channel里的数据，如果channel里没有数据的话，接收的数据是元素类型的零值
 *
 * 单向channel
 * 如果一个channel只可以接收但是不能发送，或者一个channel只能发送但不能接收，这种channel称为单向channel
 * 单向channel的声明也很简单，只需要在声明的时候带上<-操作符即可
 * onlySend := make(chan<- int)
 * onlyReceive := make(<-chan int)
 * 声明单向channel时，<-操作符的位置和上面讲到的发送和接收操作是一样的
 */
func goroutineDemo3() {
    ch := make(chan string, 5)
    ch <- "a"
    ch <- "b"
    ch <- "c"
    // cap可以获取channel的容量，len可以获取channel中元素的个数
    fmt.Printf("ch容量为:%d, 元素个数为:%d\n", cap(ch), len(ch))
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

/**
 * select+channel实现多路复用
 * 案例:启动3个goroutine进行网络资源访问，并把结果发送到3个channel中，哪个先执行完就使用哪个channel的结果
 * 这种情况下，如果我们尝试获取第一个channel的结果，程序就会被阻塞，无法判断哪个goroutine先执行完成，这个时候就需要用到多路复用了
 * Go语言中，通过select语句可以实现多路复用
 * select {
 * case i1 := <-c1:
 *     //...
 * case c2 <- i2:
 *     //...
 * default:
 *     //...
 * }
 * 整体结构和switch非常像，都有case和default，只不过select的case是一个个可以操作的channel(发送或接收)
 * 多路复用可以简单理解为在N个channel中，任意一个channel有数据产生，select都可以监听到，然后执行相应的分支接收数据并处理
 */
func selectDemo() {
    // 创建3个存放结果的channel
    firstCh := make(chan string)
    secondCh := make(chan string)
    threeCh := make(chan string)
    // 同时开启3个goroutine进行文件下载
    go func() {
        firstCh <- downloadFile("firstCh")
    }()
    go func() {
        secondCh <- downloadFile("secondCh")
    }()
    go func() {
        threeCh <- downloadFile("threeCh")
    }()
    /**
     * 开启select多路复用，哪个channel能获取到值，就说明哪个goroutine最先执行完成
     * 如果这些case中有一个可以执行，select语句会选择该case执行
     * 如果同时有多个case可以被执行，则随机选择一个，这样每个case都有平等的被执行的机会
     * 如果一个select没有任何case，那么它会一直等待下去
     */
    select {
    case filePath := <-firstCh:
        fmt.Println(filePath)
    case filePath := <-secondCh:
        fmt.Println(filePath)
    case filePath := <-threeCh:
        fmt.Println(filePath)
    }
}

func downloadFile(chanName string) string {
    min, max := 1, 10
    n := rand.Intn(max - min) + min
    fmt.Println(chanName, n)
    // 模拟文件下载
    time.Sleep(time.Second * time.Duration(n))
    return chanName + ":filePath"
}

/**
 * 在Go语言中，提倡通过通信来共享内存，而不是通过共享内存来通信
 * 其实就是提倡通过channel发送接收消息的方式进行数据传递，而不是通过修改同一个变量
 * 所以在数据流动、传递的场景中要优先使用channel，它是并发安全的，性能也不错
 */
