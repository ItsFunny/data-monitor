package dispatcher

import (
	"fmt"

	"data-monitor/data-analyser/handler"
	. "data-monitor/data-common/dispatcher"
	"data-monitor/data-common/worker"
)

type PacketAnalyserDispatcher struct {
	BaseDispatcher
}

func NewPacketAnalyserDispatcher(maxWorkerNumber, maxJobBufferNumber int) Dispatcher {
	packetAnalyserDispatcher := &PacketAnalyserDispatcher{
		BaseDispatcher: BaseDispatcher{
			MaxWorker: maxWorkerNumber,
			WorkQueue: make(chan chan interface{}, maxWorkerNumber),
			JobQueue:  make(chan interface{}, maxJobBufferNumber),
			SlefFun: func() {
				fmt.Println("[PacketAnalyserDispatcher]start")
			},
		},
	}
	for i := 0; i < maxWorkerNumber; i++ {
		packetAnalyserDispatcher.Workers[i] = &handler.PacketHandler{
			BaseWoker: worker.BaseWoker{
				WorkPool: packetAnalyserDispatcher.WorkQueue,
			},
		}
	}
	return packetAnalyserDispatcher
}
