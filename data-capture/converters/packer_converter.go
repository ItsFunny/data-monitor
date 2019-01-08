package converters

import (
	"data-monitor/data-common/models"
)

func Bytes2Packet(bytes []byte) *models.Packet {
	if len(bytes) == 0 {
		return nil
	}

	return nil
}
