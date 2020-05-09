/*
用select来调度

要实现c1和c2同时能收到数据，谁来的快收谁的，实现这种需求需要通过select
channel中无论是发送数据还是接收数据，都是阻塞式的，如果想要实现非阻塞式的，就用select+default实现

func main() {
	var c1,c2 chan int

	select{
	case n:= <- c1:	//从c1收数据
		fmt.Println("received from c1: ",n)
	case n:= <- c2:	//从c2收数据
		fmt.Println("received from c2: ",n)
	default:                         //如果c1 c2都没有数据，就走default
		fmt.Println("no value received")
	}

}


接下来实现往c1 c2发送数据
func generator() chan int{
	out := make(chan int)
	go func(){               //开完goroutine后，同时return out

	}()
	return out
}

上面写好框架后，在填充内容
func generator() chan int{
	out := make(chan int)
	go func(){               //开完goroutine后，同时return out
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i   //把i发送到out channel里
			i++
		}
	}()
	return out
}

func main() {
	var c1,c2 = generator(), generator()
	for {
		select {
		case n := <-c1:
			fmt.Println("received from c1: ",n)
		case n := <-c2:
			fmt.Println("received from c2: ",n)
		}
	}
}


在实现的复杂一些，把之前实现worker放进来

func generator() chan int{        //定义返回时一个chan of int类型
	out := make(chan int)
	go func(){               //开完goroutine后，同时return out
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i   //把i发送到out channel里
			i++
		}
	}()
	return out
}
//因为generator()函数返回的是chan类型，函数内在往这个chan里发送数据，所以当c1 := generator()时，c1只能去读取数据<-c1

func worker(id int, c chan int) {
	for n := range c{
		fmt.Printf("worker: %d received: %c",id n)
	}
}

//createWroker函数明确定义了返回类型是chan<- int，所以当w := createWorker(0)后, w<- 要收数据的
//w <-收到数据后，给go worker()这个goroutine去读取
func createWorker(id int) chan<- int{
	c := make(chan int)
	go worker(id, c)
	return c
}


func main() {
	var c1,c2 = generator(), generator() // c1 c2都是channel

	//createWorker()返回一个chan<- int，实例化后的w只能w<-, createWorker函数里面调用了worker()，这个worker函数是负责读取chan的。
	w := createWorker(0)       // w也是channel
	for {
		select {
		case n := <-c1:        
			w <- n        // 从c1读取到数据后，在发给w通道；w会调用createWorker里的go worker()函数，去读取数据
		case n := <-c2:
			w <- n
		}
	}
}


*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]

		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println(
				"queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
