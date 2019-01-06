package zmq

import "C"
import (
	"github.com/pebbe/zmq4"
	"go_dlxy/dlxy_common/constants"
)

// 监听消息,获取packet
type ZmqPakcetClient struct {
	zmqPacketClient *zmq4.Socket
	stopFlag        bool
}

func (c *ZmqPakcetClient) Init(zmqType zmq4.Type, params map[string]string) {
	ctx, err := zmq4.NewContext()
	if nil != err {
		panic(err)
	}
	c.zmqPacketClient, err = ctx.NewSocket(zmqType)
	if zmqType == zmq4.SUB {
		if topic := params["topic"]; "" != topic {
			if err = c.zmqPacketClient.SetSubscribe(topic); nil != err {
				panic(err)
			}
		} else {
			panic(constants.MissingArgumentError("subscribe topic"))
		}
	}
	if nil != err {
		panic(err)
	}
	serverAddr := params["serverAddr"]
	if serverAddr == "" {
		panic(constants.MissingArgumentError("serverAddr"))
	}
	err = c.zmqPacketClient.Connect(serverAddr)
	if nil != err {
		panic(err)
	}
}

func (c *ZmqPakcetClient) Recv() {

	msg, err := c.zmqPacketClient.RecvMessage(0)
	if nil!=err{

	}
}
