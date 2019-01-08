package worker

import (
	"data-monitor/data-common/worker"
)

type PacketWorker struct {
	worker.BaseWoker
	workQueue chan chan interface{}
	jobQueue  chan interface{}
}


func (PacketWorker) Consume(interface{}) {
	for {
		select {}
	}
}
