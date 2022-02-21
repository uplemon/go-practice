package concurrent

import (
    "fmt"
    "strconv"
    "sync"
    "time"
)

/**
 * channel为什么是并发安全的呢？是因为channel内部使用了互斥锁来保证并发的安全
 * 在Go语言中，不仅有channel这类比较易用且高级的同步机制，还有sync.Mutex、sync.WaitGroup等比较原始的同步机制
 */
func SyncDemo() {
    syncUnsafeDemo()
    syncMutexDemo()
    syncRWMutex()
    syncWaitGroup()
    syncOnceDemo()
    syncCondDemo()
}

/**
 * 资源竞争
 * 在一个goroutine中，如果分配的内存没有被其他goroutine访问，只在该goroutine中被使用，那么不存在资源竞争的问题
 * 但如果同一块内存被多个goroutine同时访问，就会产生不知道谁先访问也无法预料最后结果的情况，这就是资源竞争，这块内存可以称为共享的资源
 * 如下示例，期待的结果是1000，但很可能出现990或980等
 * 导致这种情况的核心原因是资源sum不是并发安全的，因为同时会有多个协程交叉执行sum+=i，产生不可预料的结果
 * 使用go build、go run、go test这些Go语言工具链提供的命令时，添加-race标识可以帮你检查Go语言代码是否存在资源竞争
 */
func syncUnsafeDemo()  {
    // 共享的资源
    sum := 0
    // 开启100个协程让sum+10
    for i := 0; i < 100; i++ {
        go func() {
            sum += 10
        }()
    }
    // 防止主协程提前退出
    time.Sleep(time.Second * 2)
    fmt.Println(sum)
}

/**
 * sync.Mutex
 * 互斥锁，指的是在同一时刻只有一个协程执行某段代码，其他协程都要等待该协程执行完毕后才能继续执行
 */
func syncMutexDemo() {
    var (
        sum = 0
        mutex sync.Mutex
    )
    for i :=0; i < 100; i++ {
        go func() {
            // 加锁Lock，解锁Unlock，defer语句确保锁一定会被释放
            mutex.Lock()
            defer mutex.Unlock()
            sum += 10
        }()
    }
    // 防止主协程提前退出
    time.Sleep(time.Second * 2)
    fmt.Println(sum)
}

/**
 * sync.RWMutex
 * 读写锁，该锁可以加多个读锁或者一个写锁，适用于读多写少的场景
 * sync.RWMutex比sync.Mutex性能要高，因为多个goroutine可以同时读数据，不再相互等待
 */
func syncRWMutex() {
    var (
        sum = 0
        mutex sync.RWMutex
    )
    // write
    for i := 0; i < 100; i++ {
        go func() {
            mutex.Lock()
            defer mutex.Unlock()
            sum += 10
        }()
    }
    // read
    for i := 0; i < 10; i++ {
        go func() {
            // 只获取读锁
            mutex.RLock()
            defer mutex.RUnlock()
            fmt.Println(sum)
        }()
    }
    // 防止主协程提前退出
    time.Sleep(time.Second * 2)
    fmt.Println(sum)
}

/**
 * sync.WaitGroup
 * 监听所有协程的执行，一旦全部执行完毕，程序马上退出，这样既保证了所有协程执行完毕，又可以及时退出节省时间提升性能
 * 相比使用time.Sleep()设置固定时间去等待更加精准高效
 * sync.WaitGroup的使用分为如下三步：
 * 1. 声明一个sync.WaitGroup，然后通过Add方法设置计数器的值，需要跟踪多少个协程就设置多少
 * 2. 在每个协程执行完毕时调用Done方法，让计数器减1，告诉sync.WaitGroup该协程已经执行完毕
 * 3. 最后调用Wait方法一直等待，直到计数器值为0，也就是所有跟踪的协程都执行完毕
 * 通过sync.WaitGroup可以很好地跟踪协程，在其他协程执行完毕后，主协程函数才能执行完毕
 */
func syncWaitGroup() {
    var (
        sum = 0
        mutex sync.RWMutex
        wg sync.WaitGroup
    )
    const goroutineNum = 100
    // 监听协程数，设置计数器为goroutineNum
    wg.Add(goroutineNum)
    for i := 0; i < goroutineNum; i++ {
        go func() {
            // 计数器值减1
            defer wg.Done()
            mutex.Lock()
            defer mutex.Unlock()
            sum += 10
        }()
    }
    // 一直等待直到所有协程执行完毕
    wg.Wait()
    fmt.Println(sum)
}

/**
 * sync.Once
 * 在实际的工作中可能会有这样的需求，希望代码只执行一次，Go语言为我们提供了sync.Once来保证代码只执行一次
 * sync.Once适用于创建某个对象的单例、只加载一次的资源等只执行一次的场景
 */
func syncOnceDemo() {
    var (
        once sync.Once
        wg sync.WaitGroup
    )
    const goroutineNum = 10
    wg.Add(goroutineNum)
    for i := 0; i < goroutineNum; i++ {
        go func(i int) {
            defer wg.Done()
            fmt.Println("goroutine_" + strconv.Itoa(i))
            // 保证代码执行一次
            once.Do(func() {
                fmt.Println("syncOnce")
            })
        }(i)
    }
    wg.Wait()
}

/**
 * sync.Cond
 * 在Go语言中，sync.WaitGroup用于最终完成的场景，关键点在于一定要等待所有协程都执行完毕
 * 而sync.Cond可以用于发号施令，一声令下所有协程都可以开始执行，关键点在于协程开始的时候是等待的，要等待sync.Cond唤醒才能执行
 * sync.Cond从字面意思看是条件变量，它具有阻塞协程和唤醒协程的功能，所以可以在满足一定条件的情况下唤醒协程，但条件变量只是它的一种使用场景
 * 下面以10个人赛跑为例来演示sync.Cond的用法，在示例中有1个裁判，裁判要先等这10个人准备就绪，然后一声发令枪响，这10个人就可以开始跑了
 */
func syncCondDemo() {
    cond := sync.NewCond(&sync.Mutex{})
    var wg sync.WaitGroup
    wg.Add(11)
    for i := 0; i < 10; i++ {
        go func(num int) {
            defer wg.Done()
            fmt.Printf("[%d]号已经就位\n", num)
            cond.L.Lock()
            cond.Wait() //使当前协程进入等待状态而不结束，直到在其他协程中被唤醒通知到然后继续执行完成
            cond.L.Unlock()
            fmt.Printf("[%d]号running\n", num)
        }(i)
    }
    // 等待所有goroutine都进入wait状态
    time.Sleep(time.Second * 2)
    go func() {
        defer wg.Done()
        fmt.Println("裁判已经就位，准备发令枪")
        fmt.Println("比赛开始，大家准备跑")
        cond.Broadcast() // 通知其他协程继续执行
    }()
    wg.Wait()
}
/**
 * 以上示例步骤解析:
 * 1. 通过sync.NewCond函数生成一个*sync.Cond，用于阻塞和唤醒协程
 * 2. 然后启动10个协程模拟10个人，准备就位后调用cond.Wait()方法阻塞当前协程等待发令枪响，这里需要注意的是调用cond.Wait()方法时要加锁
 * 3. time.Sleep用于等待所有人都进入wait阻塞状态，这样裁判才能调用cond.Broadcast()发号施令
 * 4. 裁判准备完毕后，就可以调用cond.Broadcast()通知所有人开始跑了
 * sync.Cond有三个方法，它们分别是：
 * 1. wait，阻塞当前协程，直到被其他协程调用Broadcast或者Signal方法唤醒，使用的时候需要加锁，使用sync.Cond中的锁即可，也就是L字段
 * 2. Signal，唤醒一个等待时间最长的协程
 * 3. Broadcast，唤醒所有等待的协程
 */
