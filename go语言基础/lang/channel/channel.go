/*

	goroutine  <-->  goroutine
	goroutine  <-->  goroutine
			  调度器

我们可以开很多的goroutine， goroutine之间的双向通道就是channel
channel是goroutine之间的交互的通道，如果发送数据后，一定要有goroutine去收它，不然会报deadlock

func chanDemo(){
	//var c chan int      c == nil
	c := make(chan int)  初始化一个channel，channel里面是int类型

	go func() {
		for {
			n := <-c             从chan收数据
			fmt.Println(n)
		}
	}()
	//向chan发送数据后，要有goroutine接收才行，不然会报deadlock
	c <- 1           向chan里面发送数据
	c <- 2

	fmt.Println(n) 
	time.Sleep(time.Millisecond)

}

channel也是一等公民，可以作为参数和返回值

//定义c为chan of int类型
func worker(id int,c chan int) {
	for {
		//n := <-c            //从chan中收数据，然后打印
		//fmt.Println(n)
		fmt.Printf("Worker: %d  received: %d", id, <-c)
	}
}

func chanDemo() {
	c := make(chan int)
	go worker(0, c)              函数式编程
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()   // 输出 Worker: 0  received: 1 
	             //      Worker: 0  received: 2
}


在继续完善channel

func worker(id int, c chan int) {
	fmt.Printf("Worker: %d  received: %c",id, <-c)
}

func chanDemo() {
	var channels [10]chan int       //给10个work，分别单独创建一个channel
	for i:=0; i<10; i++ {           //开10个work
		channels[i] = make(chan int) //初始化每个channel
		go work(i, channels[i])      //调用work，传id和chann进去; 把10个channel分发给10个worker
	
	}
	//向channel中发送数据
	for i:=0; i<10; i++ {
		channels[i] <- 'a' + i
	}

	for i:=0; i<10; i++ {
		channels[i] <- 'A' + i
	}
}

func main() {
	chanDemo()
}

channel也可以作为返回值

//这是返回channel的常见写法
func createWorker(id int) chan int {  //定义返回值为chan of int
	c := make(chan int)
	go func() {    //如果不写gorouting，如果还没有往c里发送数据，就死循环了程序就挂掉了
		for {
			fmt.Printf("worker: %d  received: %c",id, <-c)
		}
	}()
	retrun c
}

func chanDemo() {
	var channels [10]chan int
	for i:=0;i<10;i++{
		channels[i] = createWorker(i) //一共创建10个worker, 每个worker返回一个chan,然后把返回的chan存起来 
									//createWorker函数返回一个chan类型，并里面定义了goroutine来接收数据
	}

	for i:=0; i<10; i++ {           分别给10个channel发送数据
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}
返回一个channel的写法
1. 定义函数返回类型为chan  fucn test() chan int{}
2. 在函数中初始化一个chan  c:=make(chan int)
3. 然后在函数中return c  return一个chan
4. 中间的这部分应该是分发给goroutine去做的，不然后死循环了for{}，因为还没人给chan发数据，所以程序就挂掉了
5. 对channel做的事情定义的go func()里，中间这部分应该要定义goroutine。 go func(){ for {fmt.Println("...")}}


接下来在channel的定义上，在做一些修饰

func createWorker(id int) chan int{}   这个函数返回chan int, 如果这个函数比较负责，这个chan int到底怎么用很难看出来

所以要告诉外面的人要怎么用

我们可以看到:
	channels[i] = createWorker(i) 创建好channel后
	channels[i] <- 'a' + i   是往这个channel里发数据的 

所以可以改下返回值
func createWorker(id int) chan<- int{    //chan<- int   定义返回值是chan，并且是往chan中送数据的;外面的人是发数据的
	c := make(chan int)
	go func(){
		for {
			fmt.Printf("worker: %d received: %c",id,<-c) // 里面的人那就是收数据的
		}
	}
	return c
}

初始化channel的时候也要定义好类型
var channels [10]chan<- int    指定这个channel是收数据的，那么这个channel就不能往外发数据


func test() chan<- int{}  返回chan<- int， 外面的人只能收数据
a := test()
a <- 1    那么a只能收数据

func test1() <-chan int{} 返回<-chan int，外面的人只能发数据
b := test1()
n := <-b     那么b就可以发数据了



接下来讲buffer channel; 前面说了如果发数据后，没人收是不可以的,但是一旦往channel中发送数据后，就要做协程的切换是比较耗资源的
虽然协程是轻量级的，但是发完数据后一直要等人收也是比较耗资源的。所以我们可以加入缓冲区。


func budderedChannel() {
	c := make(chan int,3) //缓冲区的大小是3，可以往c这个channel发送3个数据并且不需要goroutine切换
	c <- 1
	c <- 2
	c <- 3
	//c <- 4  如果在发送一个4，程序就会报deadlock
}

buffer channel在性能上有一定的优势

现在channel有一个问题，它不知道什么时候发完
func worker(id int, c chan int) {
	for {
		fmt.Printf("worker: %d received: %c",id,<-c)
	}
}

有一种方式告诉接收方，我们的数据发完了；发送方来close的，通知接收方没有新的数据要发了
func channelClose(){
	c := make(chan int)
	go worker(0,c)  //接收chan的数据
	c <- 'a'         //给chan发数据
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)  // 关闭chan，告诉接收方已经发完了
			  // channel一旦close了，接收方还是可以收到数据的，如果chan是int类型，收到的就是0; 字符类型收到的就是空串
	time.Sleep(time.Millisecond) //close后，在收1毫秒的空串
}

func worker(id int, c chan int) {
	for {
		n, ok := <-c          //n是具体的值，ok是判断是否还有值
		if !ok {
			break
		}
		fmt.Printf("worker: %d received %d",id,n)
	}
}

上面的for循环可以可以这样写，更简单:
for n := range c {           //等c发完后，会自动跳出来
	fmt.Printf("worker: %d received %d",id,n)
}

channel是一等公民, 可以作为函数中的参数，也可以作为返回值
make(chan int,3) 给channel设置缓冲大小，buffer channel

当发送方close channel后，接收方可以有两种方式判断
1. n, ok := <-c   通过ok变量来判断
2. for n := range c {}  自动判断

不要通过共享内存来通信; 通过通信来共享内存

*/
package main

import (
	"fmt"
	"time"
)

//打印接收的数据
func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

//创建channel 然后将其返回出去  chan<- int 用来发送数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c) //因为c的类型是chan<-， 所以worker里面要去接收for n:= range c{}
	return c
}

//这里有个知识点: 
/*
1. createWorker函数返回的是一个chan<-int类型，所有当 c := createWorker(1)后，c这个channel， 应该是c <- 'xx',是往c里面发送数据
2. createWorker函数中开启了一个gorouting(go worker(id,c)), 因为c是chan<-类型，所有worker函数中要对c这个channel进行接收
3. createWorker执行后，会开一个gorouting来执行worker函数

当使用channel时，要开两个gorouting， 一个负责发送数据，一个负责接收数据
func sender(c chan) {
	go func(){}
}

func main() {
	c := make(chan int)
	sender(c) //发送的gorouting
	
	for v := range c { //接收的gorouting
		fmt.Println(v)
	}
}

*/
func chanDemo() {
	// var channels [10]chan int  定义数组，为chan of int类型
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {          
		channels[i] = createWorker(i) 
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}

小结:
当前一个函数返回的类型是一个chan时，要开一个gorouting来对这个chan进行发送或者接收的操作，然后在另外一个gorouting中实现对这个
chan的接收或者发送

func test() chan<- int{       
	c := make(chan int)
	go func(){          //开了一个gorouting去接收c的值
		for v := range c {
			fmt.Println(v)
		}
	}()
	return c
}


func main() {
	t := test()
	for i:=0;i<10;i++{    //main函数中gorouting往c里发送数据
		t <- i
	}

	time.Sleep(time.Millisecond)
}

输出: 0 1 2 3 ...9

为什么是数字是顺序的呢，因为这里只开了一个gorouting来接收，所以是顺序的，如果开10个gorouting来接收就不是顺序的了


func test() chan<- int{
	c := make(chan int)
	for i:=0;i<10;i++ {     //开10个gorouting
		go func(){
			for v := range c {
				fmt.Println(v)
			}
		}()
		}
	return c
}


func main() {
	t := test()
	for i:=0;i<10;i++{
		t <- i
	}

	time.Sleep(time.Millisecond)
}

