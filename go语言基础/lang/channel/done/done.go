/*
使用channel等待任务结束

之前channel.go中程序，有一个问题是接收方并不知道什么时候接受完毕；要保证发送数据之后，接收方要确认收到完成

//接收方
func doWork(id int, c chan int, done chan bool) { //在定义个参数 done chan bool，用来判断是否接收完成
	for n := range c {
		fmt.Printf("worker: %d received: %c", id,n)
	}
	done <- true  // 接收完成后，往done通道发送true; 通知外面事情做完了
}

//把这两个channel包装一下
type worker struct {
	in chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)  //doWork函数是接收数据，最后会往w.done通道中发一个true
	return w
}

func chanDemo(){
	var workers [10]worker       //定义worker类型的数组, worker是一个struct，包含{in done} 两个channel
	for i:=0; i<10; i++ {
		workers[i] = createWorker(i) //接收数据
	}

	//发送数据
	for i:=0; i<10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].done         在某一个位置要收这个done;收过来的数据是什么无所谓，代表接收方已经读好了
	}
	for i:=0; i<10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}
	//time.Sleep(time.Millisecond)

}
上面的操作是顺序打印的，从0-9依次打印，不是并行执行了

因为上面的代码是有两个任务，要分开来收这个done

func chanDemo() {

	//先发第一个任务
	for i, worker := range workers{
		worker.in <- 'a' + i
	}
	//收第一个任务的done通道
	for _,worker := range workers{
		<-worker.done
	}

	//在收第二个任务
	for i, worker := range workers{
		worker.in <- 'A' + i
	}
	//收第二个任务的done通道
	for _,worker := range workers{
		<-worker.done
	}

}
上面程序就是先并发的打小写，在并发的打大写




*/

/*
等待多人完成任务的方法，go语言有一个库也可以实现sync.WaitGroup

var wg sync.WaitGroup

wg.Add(20)  代表有20个任务
wg.Done   每个任务已经做完
wg.Wait  等待任务做完

使用sync.WaitGroup来完成等待多人完成任务的事情

func doWork(id int, c chan int, wg *sync.WaitGroup) {
	for n:= range c{
		fmt.Printf("Worked: %d received: %c",id,n)
		wg.Done()         代表任务做完; wg.Done是work里面调用的
	}	
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker{
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w.in, wg)
	return w

}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker

	for i:=0;i<10; i++ {
		workers[i] = createWork(i,&wg)
	}

	wg.Add(20)  20个任务

	for i,worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()  等待任务做完
}

以上就用系统提供的sync.WaitGroup来等待多人完成任务的事情



*/
package main

import (
	"fmt"
	"sync"
)

func doWork(id int,
	w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(
	id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
