package dispatcher

import (
	"data-monitor/data-capture/interfaces"
	"io"
	"log"
)

//TODO 修改dispatcher中的chan 类型为interface,consumer端做强转
// 引入BaseDispatcher,做设计模式的优化,大概就是只需要继承即可,当然当pool类型改了之后worker的类型也是需要改的

var jobQueue chan interface{}

func init() {
	// 限定大小
	jobQueue = make(chan interface{})
}

type Dispatcher interface {
	io.Closer
	Start()
}

type BaseDispatcher struct {
	workQueue chan chan interface{}
	workers   []interfaces.Worker              // 保存了哪些worker
	run       func()                           // 2019-01-06 self func ,用于处理自身的方法
	Init      func(chan chan interface{}, int) //2019-01-06 用于初始化,抽象方法,具体类具体实现
}

func (d *BaseDispatcher) Close() error {
	for _, worker := range d.workers {
		e := worker.Close()
		if nil != e {
			// TODO 更改为logrus
			log.Println(e)
		}
	}
	return nil
}

func (d *BaseDispatcher) Start() {
	d.run()
	go func() {
		for _, worker := range d.workers {
			worker.Start()
		}
	}()
	for {
		select {
		case jobBuffer := <-jobQueue:
			jobChannel := <-d.workQueue
			jobChannel <- jobBuffer
		}
	}
}

func AddJob(job interface{}) {
	jobQueue <- job
}
