package session

import (
	"fmt"
	. "go-tcp-server-test/configs"
	. "go-tcp-server-test/utils"
	"sync"
)

type TcpClientSessionManager struct {
	_networkFunctor SessionNetworkFunctors

	_sessionMap      sync.Map
	_curSessionCount int32 // 멀티스레드에서 호출된다

	sessionIndexPool *Deque
}

func NewClientSessionManager(config *NetworkConfig,
	networkFunctor SessionNetworkFunctors) *TcpClientSessionManager {
	sessionMgr := new(TcpClientSessionManager)
	sessionMgr._networkFunctor = networkFunctor
	sessionMgr._sessionMap = sync.Map{}

	sessionMgr._createSessionIndexPool(config.MaxSessionCount)

	return sessionMgr
}

func (sessionMgr *TcpClientSessionManager) _createSessionIndexPool(poolSize int) {
	sessionMgr.sessionIndexPool = NewCappedDeque(poolSize)

	for i := 0; i < poolSize; i++ {
		sessionMgr.sessionIndexPool.Append(int32(i))
	}
}

func (sessionMgr *TcpClientSessionManager) AddSession(session *TcpSession) bool {
	sessionUniqueId := session.SeqIndex
	sessionIndex := sessionMgr._allocSessionIndex()

	if sessionIndex == -1 {
		return false
	}

	_, result := sessionMgr._findSession(sessionIndex, sessionUniqueId)
	if result {
		return false
	}

	session.Index = sessionIndex
	sessionMgr._sessionMap.Store(sessionUniqueId, session)
	return true
}

func (sessionMgr *TcpClientSessionManager) _allocSessionIndex() int32 {
	index := sessionMgr.sessionIndexPool.Shift()

	if index == nil {
		return -1
	}

	return index.(int32)
}

func (sessionMgr *TcpClientSessionManager) _findSession(sessionIndex int32,
	sessionUniqueId uint64,
) (*TcpSession, bool) {
	if session, ok := sessionMgr._sessionMap.Load(sessionUniqueId); ok {
		return session.(*TcpSession), true
	}

	return nil, false
}

func (sessionMgr *TcpClientSessionManager) removeSession(sessionIndex int32, sessionUniqueId uint64) {
	sessionMgr._freeSessionIndex(sessionIndex)
	sessionMgr._sessionMap.Delete(sessionUniqueId)
}

func (sessionMgr *TcpClientSessionManager) _freeSessionIndex(sessionIndex int32) {
	sessionMgr.sessionIndexPool.Append(sessionIndex)
}

func (sessionMgr *TcpClientSessionManager) SendPacket(sessionIndex int32,
	sessionUniqueId uint64,
	sendData []byte) bool {
	session, result := sessionMgr._findSession(sessionIndex, sessionUniqueId)
	if result == false {
		fmt.Println("not found")
		return false
	}

	err := session.sendPacket(sendData)
	if err != nil {
		fmt.Println("error")
		return false
	}
	return true
}

// Send bytes to client
func (session *TcpSession) sendPacket(b []byte) error {
	_, err := session.TcpConn.Write(b)
	return err
}
