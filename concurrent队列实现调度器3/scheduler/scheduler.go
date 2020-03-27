package scheduler

import (
	"concurrent/engine"
)

type Scheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request //往workChan中发送的数据类型还是chan engine.Request
	//workChan 先要接收从worker发过来的chan
}

func (s *Scheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *Scheduler) WorkerReady(w chan engine.Request) {
	s.workChan <- w
}

func (s *Scheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workChan = make(chan chan engine.Request)

	go func() {

		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case request := <-s.requestChan:
				requestQ = append(requestQ, request)
			case workChan := <-s.workChan:
				workerQ = append(workerQ, workChan)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]

			}
		}

	}()
}
