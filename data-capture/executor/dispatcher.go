package executor

import (
	"data-monitor/data-capture/interfaces"
	"fmt"
)

//TODO 修改dispatcher中的chan 类型为interface,consumer端做强转
// 引入BaseDispatcher,做设计模式的优化,大概就是只需要继承即可,当然当pool类型改了之后worker的类型也是需要改的

var jobQueue chan interface{}

func init() {
	// 限定大小
	jobQueue = make(chan interface{})
}

type BaseDispatcher struct {
	workQueue chan chan interface{}
	workers   []interfaces.Worker // 保存了哪些worker
}

// 用于分发packet, 所有的packet都会到这来,singleton模式
// 后台会有一个线程池向这里取数据,然后进行分析
type PacketDispatcher struct {
	BaseDispatcher
}

func AddJob(job interface{}) {
	jobQueue <- job
}

func (e *PacketDispatcher) Close() error {
	for _, worker := range e.workers {
		worker.Close()
	}
	defer func() {
		if err := recover(); nil != err {
			fmt.Println(err)
		}
	}()
	return nil
}

func (e *PacketDispatcher) Start() {
	go e.run()
	for {
		select {
		case jobBuffer := <-jobQueue:
			jobChannel := <-e.workQueue
			jobChannel <- jobBuffer
		}
	}
}

func (e *PacketDispatcher) run() {
	for _, worker := range e.workers {
		worker.Start()
	}
}
