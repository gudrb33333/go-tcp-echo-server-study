package configs

import (
	"fmt"
	. "go-tcp-server-test/utils"
)

type NetworkConfig struct {
	IsTcp4Addr           bool
	BindAddress          string // ex)localhost:50000
	MaxSessionCount      int    // 최대 클라이언트 세션 수, 넉넉하게 많이 해도 괜찮음.
	MaxPacketSize        int    // 최대 패킷 크기
	MaxReceiveBufferSize int    //받은 버퍼 크기. 최소 MaxPacketSize 두배 이상.
}

func (config NetworkConfig) WriteNetworkConfig(isClientSide bool) {
	LogInfo("", 0, fmt.Sprintf("config - isClientSide: %t", isClientSide))
	LogInfo("", 0, fmt.Sprintf("config - IsTcp4Addr: %t", config.IsTcp4Addr))
	LogInfo("", 0, fmt.Sprintf("config - ClientAddress: %s", config.BindAddress))
	LogInfo("", 0, fmt.Sprintf("config - MaxSessionCount: %d", config.MaxSessionCount))
	LogInfo("", 0, fmt.Sprintf("config - MaxPacketSize: %d", config.MaxPacketSize))
	LogInfo("", 0, fmt.Sprintf("config - MaxReceiveBufferSize: %d", config.MaxReceiveBufferSize))
}
