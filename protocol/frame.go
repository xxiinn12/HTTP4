package protocol

type FrameType uint8

const (
	Request FrameType = 0x01
	Response FrameType = 0x02
)

type Frame struct {
	Type    FrameType
	Length  uint32
	Payload []byte
}
