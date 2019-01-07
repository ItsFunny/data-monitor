package monitor

import (
	"encoding/json"
	"github.com/pebbe/zmq4"

	"data-monitor/data-analyser/zmq"
	"data-monitor/data-common/models"
)

type Monitor struct {
	zmqPacketClient *zmq.ZmqPacketClient
	stopFlag        bool
}

// 需要从配置文件中获取配置
func (m *Monitor) init() {
	params := make(map[string]string)
	params["serverAddr"] = "127.0.0.1:10000"
	params["topic"] = "packet"
	m.zmqPacketClient.Init(zmq4.SUB, params)

}

func (m *Monitor) Start() {
	for {
		if !m.stopFlag {
			//
			bytes, e := m.zmqPacketClient.Recv()
			if nil != e {
				continue
			}
			packet := &models.Packet{}
			e = json.Unmarshal(bytes, packet)
			if nil != e {

			} else {
				// 处理packet
			}
		}
	}

}

func (m *Monitor) Close() error {
	if nil != m.zmqPacketClient {
		return m.zmqPacketClient.Close()
	}
	return nil
}
