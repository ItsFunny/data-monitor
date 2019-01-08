package analyser

import (
	"encoding/json"
	"fmt"

	"github.com/pebbe/zmq4"

	"data-monitor/data-analyser/log"
	"data-monitor/data-analyser/zmq"
	. "data-monitor/data-common/models"
)

type Analyser struct {
	zmqPacketClient *zmq.ZmqPacketClient
	stopFlag        bool
}

// 需要从配置文件中获取配置
func (m *Analyser) init() {
	// 从配置文件中拉取
	params := make(map[string]string)
	params["serverAddr"] = "127.0.0.1:10000"
	params["topic"] = "packet"
	m.zmqPacketClient.Init(zmq4.SUB, params)

	// 初始化日志线程池
	// 具体的从配置文件中拉取
	log.Config(20, 5)
}

func (m *Analyser) Start() {
	// 2019-01-08 17:18 FIXME big problem here,no need for this
	log.Start()
	for {
		if !m.stopFlag {
			//
			bytes, e := m.zmqPacketClient.Recv()
			if nil != e {
				continue
			}
			packet := &Packet{}
			e = json.Unmarshal(bytes, packet)
			if nil != e {
				// 往日志服务器中发送消息
				log.AddJob(NewLogDetail(WARN, fmt.Sprintf("解析json串失败,error")))
			} else {
				// 处理packet
			}
		}
	}

}

func (m *Analyser) Close() error {
	if nil != m.zmqPacketClient {
		return m.zmqPacketClient.Close()
	}
	return nil
}
