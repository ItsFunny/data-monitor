package dispatcher

import (
	"io"
	"log"

	"data-monitor/data-common/worker"
)

// TODO 修改dispatcher中的chan 类型为interface,consumer端做强转
// 引入BaseDispatcher,做设计模式的优化,大概就是只需要继承即可,当然当pool类型改了之后worker的类型也是需要改的

// var jobQueue chan interface{}

func init() {
	// 限定大小
	// jobQueue = make(chan interface{})
}

type Dispatcher interface {
	io.Closer
	Start()
	AddJob(interface{})
}

type BaseDispatcher struct {
	MaxWorker int
	WorkQueue chan chan interface{}
	// 如果只有一个dispatcher的话直接外抛成公共变量即可,但是因为不止有1个dispatcher,所以设置为变量
	JobQueue chan interface{}
	Workers  []worker.Worker // 保存了哪些worker
	SlefFun  func()          // 2019-01-06 self func ,用于处理自身的方法
	// 2019-01-06 用于初始化,抽象方法,具体类具体实现 2019-01-08 14:55 其实不需要,可以删了
	// Init      func(chan chan interface{}, int)
}

func (d *BaseDispatcher) Close() error {
	close(d.JobQueue)
	for _, worker := range d.Workers {
		e := worker.Close()
		if nil != e {
			// TODO 更改为logrus
			log.Println(e)
		}
	}
	return nil
}

func (d *BaseDispatcher) Start() {
	d.SlefFun()
	go func() {
		for _, worker := range d.Workers {
			worker.Start()
		}
	}()
	for {
		select {
		case jobBuffer := <-d.JobQueue:
			jobChannel := <-d.WorkQueue
			jobChannel <- jobBuffer
		}
	}
}

func (d *BaseDispatcher) AddJob(job interface{}) {
	d.JobQueue <- job
}
