package handler

import (
	"data-monitor/data-common/models"
	"data-monitor/data-common/utils"
	"data-monitor/data-common/worker"
)

// 2019-01-07 23:15 这个是分析具体的packet的执行者
type PacketHandler struct {
	worker.BaseWoker
}

func (h *PacketHandler) Consume(value interface{}) error {
	var (
		packet models.Packet
		ok     bool
	)
	if packet, ok = value.(models.Packet); !ok {
		return utils.ParseError(value, "packet")
	}
	return h.handlePacket(packet)
}

// 2019-01-08 17:33 流量分析
func (h *PacketHandler) handlePacket(packet models.Packet) error {

	return nil
}
