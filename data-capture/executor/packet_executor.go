package executor

import (
	"data-monitor/data-common/dispatcher"
)

// 用于分发packet, 所有的packet都会到这来,singleton模式
// 后台会有一个线程池向这里取数据,然后进行分析
type PacketDispatcher struct {
	dispatcher.BaseDispatcher
}

func (d *PacketDispatcher) Init(workPool chan chan interface{}, maxWorer int) {

}
