package zmq

import (
	"data-monitor/data-common/utils"
	"github.com/pebbe/zmq4"
)

// 监听消息,获取packet
type ZmqPacketClient struct {
	zmqPacketClient *zmq4.Socket
	stopFlag        bool
	lastRecvError   error
}

func (c *ZmqPacketClient) Init(zmqType zmq4.Type, params map[string]string) {
	ctx, err := zmq4.NewContext()
	if nil != err {
		panic(err)
	}
	c.zmqPacketClient, err = ctx.NewSocket(zmqType)
	if zmqType == zmq4.SUB {
		if topic, ok := params["topic"]; ok {
			if err = c.zmqPacketClient.SetSubscribe(topic); nil != err {
				panic(err)
			}
		} else {
			panic(utils.MissingArgument("subscribe topic"))
		}
	}
	if nil != err {
		panic(err)
	}
	serverAddr := params["serverAddr"]
	if serverAddr == "" {
		panic(utils.MissingArgument("serverAddr"))
	}
	err = c.zmqPacketClient.Connect(serverAddr)
	if nil != err {
		panic(err)
	}
}

func (c *ZmqPacketClient) Recv() ([]byte, error) {
	return c.zmqPacketClient.RecvBytes(0)
}

func (c *ZmqPacketClient) Close() error {
	if nil != c {
		return c.Close()
	}

	return nil
}
