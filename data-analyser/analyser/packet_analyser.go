package analyser

import (
	"data-monitor/data-common/dispatcher"
	"data-monitor/data-common/interfaces"
)

type PacketAnalyserDispatcher struct {
	dispatcher.BaseDispatcher
}

// 2019-01-07 23:15 这个是分析具体的packet的执行者
type PacketAnalyserWorker struct {
	interfaces.BaseWoker
}
func (a *PacketAnalyserDispatcher)consume(value interface{})error{
	// TODO  对packet进行处理
	return nil
}






