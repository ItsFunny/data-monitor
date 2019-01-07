package interfaces

import "io"

type CaptureInterface interface {
	Capture()
}

// 处理传过来的包
type PacketHandlerInterface interface {
	Handle([]byte)
}
