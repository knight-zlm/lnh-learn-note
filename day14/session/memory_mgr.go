package session

import (
	"fmt"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// 需要session对象

// MemorySessionMgr ...
type MemorySessionMgr struct {
	SessionMap map[string]Session
	rwlock     sync.RWMutex
}

// NewMemorySessionMgr ...
func NewMemorySessionMgr(id string) *MemorySessionMgr {
	s := &MemorySessionMgr{
		SessionMap: make(map[string]Session, 1024),
	}
	return s
}

// Init ...
func (m *MemorySessionMgr) Init(addr string, options ...string) error {
	return nil
}

// CreateSession ...
func (m *MemorySessionMgr) CreateSession() (Session, error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	u1 := uuid.NewV4()
	return NewMemorySession(u1.String()), nil
}

// Get ...
func (m *MemorySessionMgr) Get(sessionID string) (Session, error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	session, ok := m.SessionMap[sessionID]
	if !ok {
		err := fmt.Errorf("session_id %v is invalid", sessionID)
		return nil, err
	}
	return session, nil
}

// // Get ...
// func (m *MemorySession) Get(key string) (interface{}, error) {
// 	// 加锁
// 	m.rwlock.Lock()
// 	defer m.rwlock.Unlock()
// 	if value, ok := m.Data[key]; ok {
// 		return value, nil
// 	}
// 	err := fmt.Errorf("not find id:%s session", key)
// 	return nil, err
// }

// // Del ...
// func (m *MemorySession) Del(key string) error {
// 	m.rwlock.Lock()
// 	defer m.rwlock.Unlock()
// 	delete(m.Data, key)
// 	return nil
// }

// // Save ...
// func (m *MemorySession) Save() error {
// 	return nil
// }
