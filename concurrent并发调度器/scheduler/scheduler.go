package scheduler

import "concurrent/engine"

type Scheduler struct {
	workChan chan engine.Request
}

func (s *Scheduler) Submit(r engine.Request) {
	go func() { s.workChan <- r }()
}

func (s *Scheduler) ConfigMasterWorkChan(w chan engine.Request) {
	s.workChan = w
}
