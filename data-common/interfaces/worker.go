package interfaces

import (
	"io"
	"log"
)

type Worker interface {
	io.Closer
	consume(interface{}) error
	Start()
}
type BaseWoker struct {
	workPool chan chan interface{}
	jobQueue chan interface{}
	closed bool
	Config func()		// 启动之前的先初始化,或者说是配置
}

func (w *BaseWoker) Close() error {
	if !w.closed{
		close(w.jobQueue)
	}
	return nil
}

func (w *BaseWoker) consume(interface{})error {
	// TODO
	// 这个应该是交给子类实现的,模板模式
	panic("imple me plz ")
}

func (w *BaseWoker) Start() {
	for {
		w.workPool<-w.jobQueue
		select {
		case v:=<-w.jobQueue:
			err := w.consume(v)
			if nil!=err{
				log.Println("[Worker]consume occur error; ",err)
			}
		}
	}
}


