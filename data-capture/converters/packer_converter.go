package converters

import "data-monitor/data-capture/capture"

func Bytes2Packet(bytes []byte) *capture.Packet {
	if len(bytes) == 0 {
		return nil
	}

	return nil
}
