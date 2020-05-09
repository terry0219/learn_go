/*
https://www.cnblogs.com/f-ck-need-u/p/9986335.html

channel用于goroutines之间的通信，让它们之间可以进行数据交换。
像管道一样，一个goroutine_A向channel_A中放数据，另一个goroutine_B从channel_A取数据。

channel是指针类型的数据类型，通过make来分配内存。例如：

ch := make(chan int)
这表示创建一个channel，这个channel中只能保存int类型的数据。也就是说一端只能向此channel中放进int类型的值，另一端只能从此channel中读出int类型的值。

需要注意，chan TYPE才表示channel的类型。所以其作为参数或返回值时，需指定为xxx chan int类似的格式。

向ch这个channel放数据的操作形式为：

ch <- VALUE
从ch这个channel读数据的操作形式为：

<-ch             // 从ch中读取一个值
val = <-ch
val := <-ch      // 从ch中读取一个值并保存到val变量中
val,ok = <-ch    // 从ch读取一个值，判断是否读取成功，如果成功则保存到val变量中

其实很简单，当ch出现在<-的左边表示send，当ch出现在<-的右边表示recv。

例子:
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    go sender(ch)         // 激活了一个goroutine用于执行sender()函数，该函数每次向channel ch中发送一个字符串
    go recver(ch)         // 同时还激活了另一个goroutine用于执行recver()函数，该函数每次从channel ch中读取一个字符串
    time.Sleep(1e9)
}

func sender(ch chan string) { //向ch通道发送数据
    ch <- "malongshuai"
    ch <- "gaoxiaofang"
    ch <- "wugui"
    ch <- "tuner"
}

func recver(ch chan string) { //向ch通道接收数据
    var recv string
    for {
        recv = <-ch
        fmt.Println(recv)
    }
}

输出结果：
malongshuai
gaoxiaofang
wugui
tuner

重要知识:

1. 注意上面的recv = <-ch，当channel中没有数据可读时，recver goroutine将会阻塞在此行。
2. 由于recver中读取channel的操作放在了无限for循环中，表示recver goroutine将一直阻塞，直到从channel ch中读取到数据，读取到数据后进入下一轮循环由被阻塞在recv = <-ch上。直到main中的time.Sleep()指定的时间到了，main程序终止，所有的goroutine将全部被强制终止。
3.因为receiver要不断从channel中读取可能存在的数据，所以receiver一般都使用一个无限循环来读取channel，避免sender发送的数据被丢弃。

--------------------------------------------------------------------------------------------

channel的三种操作(send receive close)

send：表示sender端的goroutine向channel中投放数据
receive：表示receiver端的goroutine从channel中读取数据
close：表示关闭channel

关闭channel后，send操作将导致painc
关闭channel后，recv操作将返回对应类型的0值以及一个状态码false
close并非强制需要使用close(ch)来关闭channel，在某些时候可以自动被关闭
如果使用close()，建议条件允许的情况下加上defer
只在sender端上显式使用close()关闭channel。因为关闭通道意味着没有数据再需要发送

例如，判断channel是否被关闭：
	val, ok := <-counter
	if ok {
		fmt.Println(val)
	}
因为关闭通道也会让recv成功读取(只不过读取到的值为类型的空值)，使得原本阻塞在recv操作上的goroutine变得不阻塞，借此技巧可以实现goroutine的执行先后顺序

--------------------------------------------------------

channel分为两种: unbuffered channel(无缓冲区channel) 和 buffered channel(有缓冲区channel)

unbuffered channel：阻塞、同步模式
	1. sender端向channel中send一个数据，然后阻塞，直到receiver端将此数据receive
	2. receiver端一直阻塞，直到sender端向channel发送了一个数据

buffered channel：非阻塞、异步模式
	1. sender端可以向channel中send多个数据(只要channel容量未满)，容量满之前不会阻塞
	2. receiver端按照队列的方式(FIFO,先进先出)从buffered channel中按序receive其中数据

重要知识:
1. 可以认为阻塞和不阻塞是由channel控制的，无论是send还是recv操作，都是在向channel发送请求：
2. 对于unbuffered channel，sender发送一个数据，channel暂时不会向sender的请求返回ok消息，而是等到receiver准备接收channel数据了，channel才会向sender和receiver双方发送ok消息。在sender和receiver接收到ok消息之前，两者一直处于阻塞。
3. 对于buffered channel，sender每发送一个数据，只要channel容量未满，channel都会向sender的请求直接返回一个ok消息，使得sender不会阻塞，直到channel容量已满，channel不会向sender返回ok，于是sender被阻塞。对于receiver也一样，只要channel非空，receiver每次请求channel时，channel都会向其返回ok消息，直到channel为空，channel不会返回ok消息，receiver被阻塞。

-------------------------------------------------------------------------------

buffered channel的两个属性: 容量和长度

1. capacity：表示bufffered channel最多可以缓冲多少个数据
2. length：表示buffered channel当前已缓冲多少个数据
3. 创建buffered channel的方式为make(chan TYPE,CAP)

重要知识:
unbuffered channel可以认为是容量为0的buffered channel，所以每发送一个数据就被阻塞。
注意，不是容量为1的buffered channel，因为容量为1的channel，是在channel中已有一个数据，并发送第二个数据的时候才被阻塞。

换句话说，send被阻塞的时候，其实是没有发送成功的，只有被另一端读走一个数据之后才算是send成功。
对于unbuffered channel来说，这是send/recv的同步模式。而buffered channel则是在每次发送数据到通道的时候，(通道)都向发送者返回一个消息，容量未满的时候返回成功的消息，发送者因此而不会阻塞，容量已满的时候因为已满而迟迟不返回消息，使得发送者被阻塞

实际上，当向一个channel进行send的时候，先关闭了channel，再读取channel时会发现错误在send，而不是recv。它会提示向已经关闭了的channel发送数据。

小例子:
func main() {
	ch := make(chan int)
	go func(){
		ch <- 1
	}()
	close(ch)
	fmt.Println(<-ch)
}
输出结果: panic: send on closed channel

1. 创建了一个无缓冲区的ch，同步阻塞式的，有发送就一定要有接收，否则就阻塞，等main执行完就会退出
2. main本身也是一个goroutine, 加上go func(){}()，一共是有两个goroutine
3. main这个goroutine先是关闭了channel,然后在接收ch
4. go func()这个goroutine往ch里发送数据

这个流程就不对了，close(ch)这个操作应该由发送方来关，而不是由接收方关；需要把close(ch)，放到发送方来操作
func main() {
	ch := make(chan int)
	go func(){
		ch <- 1
		close(ch)
	}()
	fmt.Println(<-ch)
}

--------------------------------------------------------------------------------------------------

死锁(deadlock)

当channel的某一端(sender/receiver)期待另一端的(receiver/sender)操作，另一端正好在期待本端的操作时，也就是说两端都因为对方而使得自己当前处于阻塞状态，这时将会出现死锁问题。

更通俗地说，只要所有goroutine都被阻塞，就会出现死锁。

比如，在main函数中，它有一个默认的goroutine，如果在此goroutine中创建一个unbuffered channel，并在main goroutine中向此channel中发送数据并直接receive数据，将会出现死锁：

例子:
func chanDemo() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)

}
func main() {
	chanDemo()
}

重要知识:

在上面的示例中，向unbuffered channel中send数据的操作ch <- 1是在main goroutine中进行的，从此channel中recv的操作<-ch也是在main goroutine中进行的。
send的时候会直接阻塞main goroutine，使得recv操作无法被执行，go将探测到此问题，并报错:
fatal error: all goroutines are asleep - deadlock!

上面代码分析:
1. ch是一个无缓冲区的channel，同步阻塞模式，有发送就必须要有接收
2. main函数是一个goroutine
3. 在一个goroutine中，既发送数据ch <- 1， 又接收数据<-ch
4. 因为上面说了，这个同步阻塞模式的channel，当main这个goroutine发送数据后就阻塞了，程序根本不会再往下执行了

所以要修复这个问题，有两个方法:
1. 将ch这个channel，设置为非阻塞式异步的 make(ch chan int,1)

2. 在起一个goroutine，将send操作放在这个goroutine下执行，代码如下:

func chanDemo() {
	ch := make(chan int)
	go func(){         在新起一个goroutine，用于send数据
		ch <- 1
	}()
	fmt.Println(<-ch)	这个received操作由main goroutine完成
}

func main() {
	chanDemo()
}

*/

/*
多个管道, 输出作为输入

channel是goroutine与goroutine之间通信的基础，一边产生数据放进channel，另一边从channel读取放进来的数据。
可以借此实现多个goroutine之间的数据交换，例如goroutine_1->goroutine_2->goroutine_3，就像bash的管道一样，上一个命令的输出可以不断传递给下一个命令的输入，只不过golang借助channel可以在多个goroutine(如函数的执行)之间传，而bash是在命令之间传。

以下是一个示例，第一个函数getRandNum()用于生成随机整数，并将生成的整数放进第一个channel ch1中，第二个函数addRandNum()用于接收ch1中的数据(来自第一个函数)，将其输出，然后对接收的值加1后放进第二个channel ch2中，第三个函数printRandNum接收ch2中的数据并将其输出。

如果将函数认为是Linux的命令，则类似于下面的命令行：ch1相当于第一个管道，ch2相当于第二个管道
getRandNum | addRandNum | printRandNum

package main
import "fmt"

//起一个goroutine，对ch管道send数据，send完成后关闭close(ch)
func getRandNum(ch chan<- int) {  // ch chan<- int   ch这个管道只能send
	go func(){
		for i:=0;i<10;i++{
			ch <- i
		}
		close(ch)
	}()
}

//起一个goroutine，从in这个管道received数据后做加1操作后，在send到out的管道中，send完成后关闭管道close(out)
func addRandNum(in <-chan int , out chan<- int) { // in <-chan int  in管道只能received; out chan<- int out管道只能send
	go func(){
		for val := range in{
			val += 1
			out <- val
		}
		close(out)
	}()
}

//从out管道received数据并打印
func printRandNum(out <-chan int) {
	for val := range out {
		fmt.Println(val)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	getRandNum(ch1)
	addRandNum(ch1,ch2)
	printRandNum(ch2)
}

或者这样写也是可以的，先定义好函数，最后执行的时候调用go GetRandNum(ch1)来开启一个gorouting

func GetRandNum(n chan<- int) {
      for i:=0;i<10;i++ {
		n <- i
	}

}

func GetRandNumAdd(in <-chan int, out chan<- int) {
	for v := range in {
		v++
		out<-v
	}
}

func PrintNum(n <-chan int) {

	for v := range n {
		fmt.Println(v)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go GetRandNum(ch1)
	go GetRandNumAdd(ch1,ch2)
	go PrintNum(ch2)

        time.Sleep(time.Millisecond)
}


----------------------------------------------------------------------------------------------

select 多路监听

很多时候想要同时操作多个channel，比如从ch1、ch2读数据。
Go提供了一个select语句块，它像switch一样工作，里面放一些case语句块，用来轮询每个case语句块的send或recv情况。

select用法格式示例：

select {
    // ch1有数据时，读取到v1变量中
    case v1 := <-ch1:
        ...
    // ch2有数据时，读取到v2变量中
    case v2 := <-ch2:
        ...
    // 所有case都不满足条件时，执行default
    default:
        ...
}
defalut语句是可选的，不允许fall through行为，但允许case语句块为空块。select会被return、break关键字中断：
return是退出整个函数，break是退出当前select。

重要知识:

select的行为模式主要是对channel是否可读进行轮询，但也可以用来向channel发送数据。它的行为如下：

1. 如果所有的case语句块评估时都被阻塞，则阻塞直到某个语句块可以被处理
2. 如果多个case同时满足条件，则随机选择一个进行处理，对于这一次的选择，其它的case都不会被阻塞，而是处理完被选中的case后进入下一轮select(如果select在循环中)或者结束select(如果select不在循环中或循环次数结束)
3. 如果存在default且其它case都不满足条件，则执行default。所以default必须要可执行而不能阻塞

select语句是在某一个goroutine中运行的，就不难理解只有所有case都不满足条件时，select所在goroutine才会被阻塞，只要有一个case满足条件，本次select就不会出现阻塞的情况。
需要注意的是，如果在select中执行send操作，则可能会永远被send阻塞。所以，在使用send的时候，应该也使用defalut语句块，保证send不会被阻塞。如果没有default，或者能确保select不阻塞的语句块，则迟早会被send阻塞。在后文有一个select中send永久阻塞的分析：双层channel的一个示例。

一般来说，select会放在一个无限循环语句中，一直轮询channel的可读事件。

下面是一个示例，pump1()和pump2()都用于产生数据(一个产生偶数，一个产生奇数)，并将数据分别放进ch1和ch2两个通道，rece()则从ch1和ch2中读取数据。
然后在无限循环中使用select轮询这两个通道是否可读，最后main goroutine在1分钟后强制中断所有goroutine


*/

package main

import "fmt"
import "time"

//
func pump1(ch chan int) {
	for i := 0; i < 1000; i++ {
		if i%2 == 1 {
			ch <- i
		}
	}
	close(ch) //send主动close后，如果没有数据后, recv接收到到的是int类型的默认值0
}

func pump2(ch chan int) {
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			ch <- i
		}
	}
	//close(ch)
}

//received select
func rece(ch1, ch2 chan int) {
	for {
		select {
		case n := <-ch1:
			fmt.Printf("channel: %s count: %d\n", "ch1", n)
		case n := <-ch2:
			fmt.Printf("channel: %s count: %d\n", "ch2", n)
		default: //如果send方，主动关闭了close就不会走default分支; 如果send方没有主动关闭，就走default分支
			fmt.Println("no value")
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//3个goroutine并行执行
	go pump1(ch1)
	go pump2(ch2)
	go rece(ch1, ch2)

	time.Sleep(time.Minute) //main函数sleep 1分钟后在退出，main goroutine退出后，其他的goroutine也会被杀掉
}
