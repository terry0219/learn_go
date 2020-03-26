package engine

import "fmt"

type Request struct{
	Url string
}

type ConcurrentEngine struct {
	Scheduler Scheduler
}

type Scheduler interface {
	Submit(Request)
	Run()
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	c.Scheduler.Run()

	for _,v:= range seeds {
		//send seeds to scheduler
		for {
			fmt.Println("发送数据-->")
			c.Scheduler.Submit(v)
		}
		
	}



}