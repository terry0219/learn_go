/*
并发调度器

架构：


			request									channel
engine   ---------------->  scheduler(gorouting)  -----------------> 多个worker(gorouting)
			(函数调用传递)


scheduler通过in channel发送数据给worker, worker处理完后将数据发送到out channel.
多个worker共用一个in channel; scheduler和worker共用in通道; worker和engine共用out通道.
engine通过循环接收out channel，在传给scheduer.

*/


package main

import (
	"concurrent/engine"
	"concurrent/scheduler"
)

func main() {
	e := engine.SimpleEngine{
		Scheduler: &scheduler.Scheduler{},
	}

	e.Run(engine.Request{Url: "www.123.com"})

}
