package scheduler

import (
	"concurrent/engine"
	"fmt"
)

type Scheduler struct {
	requestChan chan engine.Request
}

func (s *Scheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *Scheduler) Run() {
	s.requestChan = make(chan engine.Request)

	go func() {

		var requestQ []engine.Request

		for {

			request := <-s.requestChan
			fmt.Println("接收数据")
			requestQ = append(requestQ, request)
			if len(requestQ) > 0 {
				activeRequest := requestQ[0]
				fmt.Println(activeRequest)
			}
		}

	}()
}
