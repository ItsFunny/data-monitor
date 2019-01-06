package monitor

import (
	"data-monitor/data-monitor/zmq"
	"github.com/pebbe/zmq4"
)

type Monitor struct {
	zmqPacketClient zmq.ZmqPakcetClient
}

// 需要从配置文件中获取配置
func (m *Monitor) init() {
	params := make(map[string]string)
	params["serverAddr"] = "127.0.0.1:10000"
	params["topic"] = "packet"
	m.zmqPacketClient.Init(zmq4.SUB, params)

}

func (m *Monitor) Start() {
}

func (m *Monitor) Close() error {
	if nil != m.zmqClient {
		return m.zmqClient.Close()
	}
	return nil
}
