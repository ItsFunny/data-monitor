package analyser

import "data-monitor/data-common/dispatcher"

type PacketAnalyserDispatcher struct {
	dispatcher.BaseDispatcher
}

type PacketAnalyser struct {
}

func (PacketAnalyser) Close() error {
	return nil
}

func (PacketAnalyser) Consume(interface{}) {

}

func (PacketAnalyser) Start() {

}
