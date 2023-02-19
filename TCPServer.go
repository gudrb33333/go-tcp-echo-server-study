package main

import (
	"fmt"
	. "go-tcp-server-test/configs"
	. "go-tcp-server-test/session"
	. "go-tcp-server-test/utils"
	"log"
	"net"
	"strconv"
	"strings"
	"sync/atomic"
)

type TCPServer struct {
	ServerIndex int
	IP          string
	Port        int
}

func createServer(netConfig NetworkConfig) {
	LogInfo("", 0, "CreateServer !!!")

	var server TCPServer

	if server.setIPAddress(netConfig.BindAddress) == false {
		LogError("", 0, "fail. server address")
		return
	}

	networkFunctor := SessionNetworkFunctors{}
	networkFunctor.OnConnect = server.OnConnect
	networkFunctor.OnReceive = server.OnReceive
	networkFunctor.OnReceiveBufferedData = nil
	networkFunctor.OnClose = server.OnClose
	networkFunctor.PacketTotalSizeFunc = PacketTotalSize
	networkFunctor.PacketHeaderSize = PACKET_HEADER_SIZE
	networkFunctor.IsClientSession = true

	defer PrintPanicStack()

	// 아래 함수가 호출되면 무한 대기에 들어간다
	TcpSessionManager = NewClientSessionManager(&netConfig, networkFunctor)
	startTcpserverBlock(&netConfig, networkFunctor)
}

func (server *TCPServer) setIPAddress(ipAddress string) bool {
	results := strings.Split(ipAddress, ":")
	if len(results) != 2 {
		return false
	}

	server.IP = results[0]
	server.Port, _ = strconv.Atoi(results[1])

	return true
}

func (server *TCPServer) OnConnect(sessionIndex int32, sessionUniqueID uint64) {
	LogInfo("", 0, fmt.Sprintf("[OnConnect] sessionIndex: %d", sessionIndex))
	LogInfo("", 0, fmt.Sprintf("[OnConnect] sessionUniqueID: %d", sessionUniqueID))
}

func (server *TCPServer) OnReceive(sessionIndex int32, sessionUniqueID uint64, data []byte) bool {
	sendToClient(sessionIndex, sessionUniqueID, data)
	LogInfo("", 0, fmt.Sprintf("[OnReceive] message %s", data))
	return true
}

func sendToClient(sessionIndex int32, sessionUniqueID uint64, data []byte) bool {
	result := TcpSessionManager.SendPacket(sessionIndex, sessionUniqueID, data)
	return result
}

func (server *TCPServer) OnClose(sessionIndex int32, sessionUniqueID uint64) {
	LogInfo("", 0, fmt.Sprintf("[OnConnect] sessionIndex: %d", sessionIndex))
	LogInfo("", 0, fmt.Sprintf("[OnConnect] sessionUniqueID: %d", sessionUniqueID))
}

func startTcpserverBlock(config *NetworkConfig, networkFunctor SessionNetworkFunctors) {
	defer PrintPanicStack()
	LogInfo("", 0, "tcpServerStart - Start")

	var err error
	tcpAddr, _ := net.ResolveTCPAddr("tcp", config.BindAddress)
	_mClientListener, err = net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer _mClientListener.Close()

	log.Println("Server Listen ...")

	for {
		conn, _ := _mClientListener.Accept()
		client := &TcpSession{
			SeqIndex:       SeqNumIncrement(),
			TcpConn:        conn,
			NetworkFunctor: networkFunctor,
		}

		TcpSessionManager.AddSession(client)

		go client.HandleTcpRead(networkFunctor)
	}

	LogInfo("", 0, "tcpServerStart - End")
}

func SeqNumIncrement() uint64 {
	newValue := atomic.AddUint64(&_seqNumber, 1)
	return newValue
}

var _mClientListener *net.TCPListener
var _seqNumber uint64 // 절대 이것을 바로 사용하면 안 된다!!!
