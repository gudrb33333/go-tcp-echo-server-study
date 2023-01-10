package utils

import (
	"encoding/binary"
)

func PacketTotalSize(data []byte) int16 {
	totalsize := binary.LittleEndian.Uint16(data)
	return int16(totalsize)
}
