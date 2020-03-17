package engine

import "fmt"

type SimpleEngine struct {
	Scheduler Scheduler
}

type Request struct {
	Url string
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkChan(chan Request)
}

func (s *SimpleEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan Request)

	s.Scheduler.ConfigMasterWorkChan(in)

	for i := 0; i < 10; i++ {
		CreateWork(i, in, out)
	}

	for _, v := range seeds {
		s.Scheduler.Submit(v)
	}

	for {
		request := <-out
		s.Scheduler.Submit(request)
	}
}

func CreateWork(id int, in chan Request, out chan Request) {
	go func() {
		for {
			request1 := <-in
			fmt.Println(id, request1)
			result := worker(request1)
			out <- result

		}
	}()
}

func worker(url Request) Request {
	//fmt.Println("working")
	return url
}
