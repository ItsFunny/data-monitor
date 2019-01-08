package worker

import (
	"io"
	"log"
)

type Worker interface {
	io.Closer
	Start()
}

// 2019-01-08 16:07 面向接口编程,干脆把这个给单独抽出
type Consumer interface {
	Consume(interface{}) error
}

type BaseWoker struct {
	WorkPool chan chan interface{}
	jobQueue chan interface{}
	closed   bool
	Config   func() // 启动之前的先初始化,或者说是配置
	// 抽象方法 ,2019-01-08 15:54 将这个抽离为单独的方法,因为不拆分的话会变成一坨代码
	//Consume  func(interface{}) error
}

// 2019-01-08 16:11 交给子类去实现
func (w *BaseWoker) Consume(interface{}) error {
	panic("implement me")
}

func (w *BaseWoker) Close() error {
	if !w.closed {
		close(w.jobQueue)
	}
	return nil
}

func (w *BaseWoker) Start() {
	for {
		w.WorkPool <- w.jobQueue
		select {
		case v := <-w.jobQueue:
			// 2019-01-08 15:30 返回的都是packet数组 ,如果返回的是packet数组的话每个worker的
			// jobQueue的buffer就得设置的小点了,如果是byte数组的话则大点就行
			err := w.Consume(v)
			if nil != err {
				log.Println("[Worker]consume occur error; ", err)
			}
		}
	}
}
