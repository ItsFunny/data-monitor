package capture

import (
	"data-monitor/data-capture/executor"
	"data-monitor/data-capture/interfaces"
	"data-monitor/data-common/constants"
	"data-monitor/data-common/models"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"io"
	"net"
	"time"
)

type CaptureDevice struct {
	// 1. 对应的设备的名称
	ifaceName string
	// 2. 对应的设备句柄
	handler *pcap.Handle
	// 3. 数据缓存
	metaPackets []*models.Packet
	// 4. 抓包的停止标志
	stopFlag bool
	// 5. 上一个错误
	previousError error
	// 6. 处理packet的对象
	packetHandler interfaces.PacketHandlerInterface
	// 7. 上次清空缓存时间,模仿延迟确认的做法,没200ms清空缓存一次
	lastFlushTime time.Duration
	// 8. 记录总共进行了多少次抓包
	pollTimes int64
}

func (c *CaptureDevice) Close() error {
	if nil != c.handler {
		c.handler.Close()
	}
	return nil
}

func StartCapture(configPath string) {
	config := LoadConfig(configPath)
	for id, iface := range config.ifaces {
		startWithConcreateInterface(id, iface)
	}
}

func startWithConcreateInterface(id int, ifaceName string) (*CaptureDevice, error) {

	_, e := net.InterfaceByName(ifaceName)
	if nil != e {
		return nil, constants.MissingError("interface", ifaceName)
	}
	// 从配置文件中拉取
	handle, e := pcap.OpenLive(ifaceName, 2048, false, time.Duration(time.Second*10))
	if nil != e {
		return nil, constants.PCAP_OPENLIVE_ERROR
	}
	capture := &CaptureDevice{
		ifaceName:   ifaceName,
		handler:     handle,
		metaPackets: make([]*models.Packet, 0, 2048),
	}
	go capture.capture()
	return capture, nil
}
func (c *CaptureDevice) capture() {
	source := gopacket.NewPacketSource(c.handler, c.handler.LinkType())
	for !c.stopFlag {
		packet, e := source.NextPacket()
		c.pollTimes++
		if nil != e {
			if e == pcap.NextErrorTimeoutExpired {

			} else if e == pcap.NextErrorReadError {

			} else if e == io.EOF {

			} else {

			}
		}
		// TODO 20190106 这里需要从池中去获取,而不是通过直接new的方式
		p := &models.Packet{Packet: packet, CaptureTime: time.Duration(time.Now().UnixNano())}
		//go c.handlePacket(p)
		// 2019-01-06 这里不额外的起线程处理了,会oom
		c.handlePacket(p)
		// 抓包,抓如下的包
		// 1.正常的包(乱序是正常的,不考虑)
		// 2.异常流量的包:
		// dup ack+retransaction+sync-flood 的包
		// 衍生为多种策略
	}
}

func (c *CaptureDevice) handlePacket(packet *models.Packet) {
	// 2019-01-05 为了避免oom,所以这里最好是都通过
	// 2019-01-06 这里不通过没次得到一个包就入队去分析,而是当达到一定数量之后再分析
	c.metaPackets = append(c.metaPackets, packet)
	if len(c.metaPackets) >= cap(c.metaPackets) || packet.CaptureTime-c.lastFlushTime > 200*time.Millisecond {
		c.Flush()
	}
}
func (c *CaptureDevice) Flush() {
	if len(c.metaPackets) == 0 {
		return
	}
	c.lastFlushTime = time.Duration(time.Now().UnixNano())
	executor.AddJob(c.metaPackets)
}
