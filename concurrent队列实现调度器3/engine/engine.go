package engine

import "fmt"

type Request struct {
	Url string
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	ItemChan chan interface{} //增加的部分
}

type Scheduler interface {
	Submit(Request)
	Run()
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan Request)

	c.Scheduler.Run() //开启了1个gorouting，启动scheduler

	for _, v := range seeds {
		//send seeds to scheduler
		fmt.Println("第一次发送数据-->")
		c.Scheduler.Submit(v)
	}

	//开10个gorouting 开启worker
	for i := 0; i < 10; i++ {
		go createWorker(out, c.Scheduler)
	}

	for {
		request := <-out

		go func() { c.ItemChan <- request }() //增加

		c.Scheduler.Submit(request)
	}

}

func createWorker(out chan Request, s Scheduler) {
	in := make(chan Request)
	for {
		s.WorkerReady(in)
		request := <-in
		result := worker(request)
		out <- result

	}
}
func worker(Request) Request {
	fmt.Println("worker 数据处理")
	return Request{}
}
