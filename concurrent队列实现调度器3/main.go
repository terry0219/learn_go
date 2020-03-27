package main

import (
	"concurrent/engine"
	"concurrent/persist"
	"concurrent/scheduler"
	"time"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.Scheduler{},
		ItemChan: persist.ItemSaver(), //新增
	}

	e.Run(engine.Request{
		Url: "www.123.com",
	})

	time.Sleep(time.Minute * 1)
}
