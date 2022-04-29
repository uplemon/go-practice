package concurrent

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextDemo() {
	watchDogDemo()
	contextWatchDogDemo()
	contextWatchDogDemo1()
	contextValueDemo()
}

/**
 * 一个协程启动后，大部分情况需要协程里的代码执行完毕，然后协程会自行退出，
 * 如果要协程提前退出怎么办呢？可以通过select+channel的方式来解决。
 * 通过channel发送指令让监控程序停止，进而达到协程退出的目的。
 */
func watchDogDemo() {
	var wg sync.WaitGroup
	stopCh := make(chan bool) // 用来停止监控程序
	wg.Add(1)
	go func() {
		defer wg.Done()
		watchDog(stopCh, "[monitor]")
	}()
	time.Sleep(time.Second * 5) // 先让监控程序监控5秒
	stopCh <- true              // 发送停止指令
	wg.Wait()
}
func watchDog(stopCh chan bool, name string) {
	// 开启for select循环，一直后台监控
	for {
		select {
		case <-stopCh:
			fmt.Printf("%s停止指令已收到，马上停止\n", name)
			return
		default:
			fmt.Printf("%s正在监控...\n", name)
			// ...
		}
		time.Sleep(time.Second * 1)
	}
}

/**
 * 通过select+channel让协程退出的方式比较优雅，但如果要同时取消多个协程该如何处理？
 * 这时select+channel局限就凸显出来了，即使定义多个channel解决问题，代码逻辑也会非常复杂不好维护。
 * 要解决这种复杂的协程问题，必须有一种可以跟踪协程的方案，只有跟踪到每个协程才能更好的控制它们，Go语言标准库提供了Context用来解决这类问题。
 */
func contextWatchDogDemo() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		contextWatchDog(ctx, "[monitor]")
	}()
	time.Sleep(time.Second * 5) // 先让监控程序监控5秒
	stop()                      // 发送停止指令
	wg.Wait()
}
func contextWatchDog(ctx context.Context, name string) {
	// 开启for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s停止指令已收到，马上停止\n", name)
			return
		default:
			fmt.Printf("%s正在监控...\n", name)
			// ...
		}
		time.Sleep(time.Second * 1)
	}
}

/**
 * 相比select+channel的方案，Context方案主要有4个改动点：
 * 1. watchDog的stopCh参数换成了ctx，类型为context.Context。
 * 2. 原来的case <-stopCh改为case <-ctx.Done()，用于判断是否停止。
 * 3. 使用context.WithCancel(context.Background())函数生成一个可以取消的Context，用于发送停止指令，
 *    这里的context.Background()用于生成一个空Context，一般作为整个Context树的根节点。
 * 4. 原来的stopCh <- true停止指令，改为context.WithCancel函数返回的取消函数stop()。
 */

/**
 * 什么是Context
 * 一个任务会有很多个协程协作完成，一次HTTP请求也会触发很多个协程的启动，
 * 而这些协程有可能会启动更多的子协程，并且无法预知有多少层协程、每一层有多少个协程。
 * 如果因为某些原因导致任务终止了，HTTP请求取消了，那么它们启动的协程怎么办？该如何取消呢？
 * 因为取消这些协程可以节约内存，提升性能，同时避免不可预料的Bug。
 * Context就是用来简化解决这些问题的，并且是并发安全的。
 * Context是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。
 * 一旦取消指令下达，那么被Context跟踪的这些协程都会收到取消信号，就可以做清理和退出操作了。
 * Context接口有四个方法，如下所示：
 * type Context interface {
 *     Deadline() (deadline time.Time, ok bool)
 *     Done() <-chan struct{}
 *     Err() error
 *     Value(key interface{}) interface{}
 * }
 * 1. Deadline方法可以获取设置的截止时间，第一个返回值deadline是截止时间，到了这个时间点，Context会自动发起取消请求，第二个返回值ok代表是否设置了截止时间。
 * 2. Done方法返回一个只读的channel，类型为struct{}。在协程中，如果该方法返回的chan可以读取，则意味着Context已经发起了取消信号。通过Done方法收到这个信号后，就可以做清理操作，退出协程释放资源。
 * 3. Err方法返回取消的错误原因，即因为什么原因Context被取消。
 * 4. Value方法获取该Context上绑定的值，是一个键值对，所以要通过一个key才可以获取对应的值。
 * Context接口的四个方法中最常用的就是Done方法，它返回一个只读的channel用于接收取消信号。当Context取消的时候会关闭这个只读channel，也就等于发出了取消信号。
 */

/**
 * Context树
 * 我们不需要自己实现Context接口，Go语言提供了函数可以帮助我们生成不同的Context，通过这些函数可以生成一颗Context树，
 * 这样Context才可以关联起来，父Context发出取消信号的时候，子Context也会发出，这样就可以控制不同层级的协程退出。
 * 从使用功能上分，有四种实现好的Context：
 * 1. 空Context：不可取消，没有截止时间，主要用于Context树的根节点。
 * 2. 可取消的Context：用于发出取消信号，当取消的时候，它的子Context也会取消。
 * 3. 可定时取消的Context：多了一个定时的功能。
 * 4. 值Context：用于存储一个key-value键值对。
 * 在整个Context树中，最顶部的是空Context，它作为整棵Context树的根节点，在Go语言中，可以通过context.Background()获取一个根节点Context。
 * 有了根节点Context后，要生成Context树，Go语言提供了如下函数：
 * 1. WithCancel(parent Context)：生成一个可取消的Context。
 * 2. WithDeadline(parent Context, d time.Time)：生成一个可定时取消的Context，参数d为定时取消的具体时间。
 * 3. WithTimeout(parent Context, timeout time.Duration)：生成一个可超时取消的Context，参数timeout用于设置多久后取消。
 * 4. WithValue(parent Context, key, val interface{})：生成一个可携带key-value键值对的Context。
 * 以上四个生成Context的函数中，前三个都属于可取消的Context，它们是一类函数，最后一个是值Context，用于存储一个key-value键值对。
 */

/**
 * 使用Context取消多个协程
 * 取消多个协程也比较简单，把Context作为参数传递给协程即可。
 * 如下示例一个Context同时控制三个协程，一旦Context发出取消信号，这三个协程都会取消退出。
 * 如果一个Context有子Context，当该Context取消时，该节点下的所有子Context都会被取消。
 */
func contextWatchDogDemo1() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())
	wg.Add(3)
	go func() {
		defer wg.Done()
		contextWatchDog1(ctx, "[monitor_1]")
	}()
	go func() {
		defer wg.Done()
		contextWatchDog1(ctx, "[monitor_2]")
	}()
	go func() {
		defer wg.Done()
		contextWatchDog1(ctx, "[monitor_3]")
	}()
	time.Sleep(time.Second * 5) // 先让监控程序监控5秒
	stop()                      // 发送停止指令
	wg.Wait()
}
func contextWatchDog1(ctx context.Context, name string) {
	// 开启for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s停止指令已收到，马上停止\n", name)
			return
		default:
			fmt.Printf("%s正在监控...\n", name)
			// ...
		}
		time.Sleep(time.Second * 1)
	}
}

/**
 * Context传值
 * 通过Context还可以传递共享的数据，比如将一个请求处理或调用过程串起来，就可以在函数调用的时候传递context。
 */
func contextValueDemo() {
	ctx := context.Background()
	// ctx是一个空context
	process(ctx)
	// 通过context.WithValue函数为context赋值
	ctx = context.WithValue(ctx, "traceId", "t_1156813585")
	process(ctx)
}
func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("traceId:%s\n", traceId)
	} else {
		fmt.Printf("no traceId\n")
	}
}

/**
 * Context使用原则
 * Context是一种非常好的工具，使用它可以很方便地控制取消多个协程。在Go语言标准库中也使用了它们，比如net/http中使用Context取消网络的请求。
 * 要更好地使用Context，有一些使用原则需要尽可能地遵守。
 * 1. Context不要放在结构体中，要以参数的方式传递。
 * 2. Context作为函数的参数时，要放在第一位，也就是第一个参数。
 * 3. 要使用context.Background函数生成根节点的Context，也就是最顶层的Context。
 * 4. Context传值要传递必须的值，而且要尽可能地少，不要什么都传。
 * 5. Context多协程安全，可以在多个协程中放心使用。
 * 以上原则是规范类的，Go语言的编译器并不会做这些检查，要靠自己遵守。
 *
 * 总结
 * Context通过With系列函数生成Context树，把相关的Context关联起来，这样就可以统一进行控制。
 * 一声令下，关联的Context都会发出取消信号，使用这些Context的协程就可以收到取消信号，然后清理退出。
 * 在定义函数的时候，如果想让外部给定义的函数发取消信号，就可以为这个函数增加一个Context参数，让外部的调用者可以通过Context进行控制，比如下载一个文件超时退出的需求。
 */
