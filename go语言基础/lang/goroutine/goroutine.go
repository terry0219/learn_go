/*
go语言并发编程

go func() {}  在函数前加一个go,就会并发的执行


在操作系统里，开10个或者100个线程都没问题，但是开1000个线程的话，就比较困难了

协程Coroutine
1. 轻量级的线程
2. 非抢占式多任务处理，由协程主动交出控制权； (线程会经常被操作系统切换，所以线程是抢占式多任务处理)
3. 编译器/解释器/虚拟机层面的多任务(协程不是操作系统层面的多任务)
4. 多个协程可能在一个或多个线程上运行, 这是由调度器决定的


func main(){
	a := [10]int
	for i:=0; i<10;i++ {
		go func(i int) {
			for {
				a[i]++   数组里面每个对应的位置做加1操作
				//runtime.Gosched() 手动交出控制权，让其他协程有机会执行
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

上面这段代码for{a[i]++} 只是一段指令，不会在协程之间做切换。这样就会被一个协程抢掉，这个协程如果不主动交出控制权，那么就会一直在当前的协程中
执行后，这段程序就是一直在运行不会退出，查看top，会有一个__goroutine的进程。

程序退不出来的原因是goroutine没有机会交出控制权，就会死在里面。main函数它自己也是一个goroutine，其他goroutine是main函数开出来的。
因为没人交出控制权，所以time.Sleep根本没机会执行


交出控制权
1.fmt.Println()是一个io操作，io操作会有一个协程之间的切换，可以交出控制权
2.runtime.Gosched() 手动交出控制权，让别人也有机会运行；调度器可以调度到其他协程中执行。一般比较少用到，有其他方式交出控制权

*/
package main

import (
	"fmt"
	"time"
)

func main() {
	//主程序在跑(for i:=0;)，它并发的开了一个程序go func(){}，相当于开了一个线程，其实是协程
	for i := 0; i < 1000; i++ { //有1000个并发执行
		//go func(i int) 这个i变量，和上面的for i:=0里的i是不一样的
		go func(i int) {  //go后面可以跟匿名函数，也可以把函数单独提出去,在调用go test() {}
			for {
				fmt.Printf("Hello from "+   //并发执行的内容
					"goroutine %d\n", i)  //这个i如果引用外面的for循环里的i不太安全，所以go func(i int)，把i传进函数里
			}
		}(i)  //调用函数
	}
	//不要让main函数马上退出；因为main函数一旦退出了，所有的goroutine就会被杀掉, goroutine里的内容还没来的及打印就被干掉了。
	//main一退出，程序就结束了
	time.Sleep(time.Minute) //在1毫秒时间内，这1000个人同时打印hello from go routine
	//如果不加time.Sleep的话，程序什么都没有打印;因为是并发执行的，main函数和for{fmt.Println("hello from")}函数是并发在执行
	//for{fmt.Println("hello from")}还没来的及打印内容，main函数中的for循环就已经跑完退出了
}

/*

如果go func(){} 里面不定义形参，a[i]++中的i就是引用了外面for循环的i变量，执行会报错index out of range

-race检测数据访问的冲突
go run -race goroutine.go


func main() {
	var a [10]int

	for i:=0; i<10; i++ {
		go func(){
			for{
				a[i]++               这里的i是引用函数外面的i，这是一个闭包；
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Millisecond)  当程序跳到这里的时候，最后一次i就等于10；那么a[i]++ 这时候就会引用a[10],
									所以就会报index out of range。因此要把i固定下来，把i传进去
	fmt.Println(a)
}

这样就把i传进go func(){}里
for i:=0;i<10;i++ {
	go func(ii int){
		for{
			a[ii]++
			runtime.Gosched()
		}
	}(i)
}


go语言的调度器

普通函数与协程的区别

普通函数: 有一个线程，它有一个main函数，在main函数中调用了doWork函数，等doWork函数执行完成后，在把控制权叫给main函数
			main --> doWork      是一个单向的； main和doWork运行在一个线程中

协程:  main <--> doWork  是一个双向的通道, main和doWork之间的数据、控制权可以双向的流通; main和doWork可能运行在一个或多个线程中。

python中的协程
1. 使用yield关键字实现协程
2. python3.5以后加入了async def 对协程原生支持


go语言的调度器负责调度协程, 一个线程中可能只有1个goroutine、或者2个goroutine，或者更多的goroutine

1. 任何函数只要加上go就能送给调度器运行
2. 调度器在合适的点进行切换
3. 使用-race来检测数据访问冲突

goroutine可能的切换点
1. I/O  select       print就是一个io操作
2. channel
3. 等待锁
4. 函数调用(有时)
5. runtime.Gosched()

虽然开了1000个协程并发运行，实际上可能系统上只开了几个线程在执行




*/



package main
import (
	"fmt"
	"time"
)

func main() {
    var a [10]int

    for i:=0;i<10;i++{
        go func(ii int){
	        //fmt.Println(ii,a[ii])
                for {
                 a[ii]++
        	 fmt.Println("gorouting: ",ii,a[ii])
                }
        }(i)
    }

    time.Sleep(time.Minute)
    //time.Sleep(time.Millisecond)
}

上面这段代码，开了10个gorouting和main函数中的gorouting

1. 10个gorouting是一个for{}无限循环，但是main函数中的gorouting定义了time.Sleep，所以当main函数中的sleep到时间后，其他10个gorouting都会
退出，并不会无限循环

2. fmt.Println("gorouting: ",ii,a[ii]) 打印当前执行的gorouting编号以及这个gorouting中对a[ii]的操作后的值


-------------------------------------------------------------------------

首先需要理解的是main函数中的gorouting和用for循环开的10个gorouting是并行执行的

package main
import (
	"fmt"
	"time"
        _ "runtime"
)

func main() {
    var a [10]int

    for i:=0;i<10;i++{
        go func(ii int){
	        //fmt.Println(ii,a[ii])
                for {
                   a[ii]++
                   //runtime.Gosched()
        	   //fmt.Println("gorouting: ",ii,a[ii])
                }
        }(i)
    }

    fmt.Println("main gorouting")
    time.Sleep(time.Millisecond)
    fmt.Println(a)
}

这个程序打印"main gorouting"后，程序就一直在死循环了。因为go func(){} 抢到控制权后，里面的操作for{a[ii]++}是一个表达式，不会交出控制权，
所以main函数中的gorouting(time.Millisecond)执行不了。

原因是: 程序退不出来的原因是goroutine没有机会交出控制权，就会死在里面。main函数它自己也是一个goroutine，其他goroutine是main函数开出来的。
因为没人交出控制权，所以time.Sleep根本没机会执行。


