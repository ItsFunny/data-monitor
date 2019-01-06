package models

import (
	"github.com/google/gopacket"
	"time"
)

type Packet struct {
	//1.
	Packet gopacket.Packet
	// 2. 抓到的时间
	CaptureTime time.Duration
}
