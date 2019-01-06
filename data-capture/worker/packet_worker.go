package worker

import "fmt"

type BaseWorker struct {
	stopChan chan struct{}
}

type PacketWorker struct {
	BaseWorker
	workQueue chan chan interface{}
	jobQueue  chan interface{}
}

func (w *PacketWorker) Start() {
	for {
		w.workQueue <- w.jobQueue
		select {
		case job := <-w.jobQueue:
			w.Consume(job)
		case <-w.stopChan:
			// TODO
			fmt.Println("stop")
		}
	}
}

func (w *PacketWorker) Close() error {
	close(w.jobQueue)
	return nil
}

func (PacketWorker) Consume(interface{}) {
	for {
		select {}
	}
}
