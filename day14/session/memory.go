package session

import (
	"fmt"
	"sync"
)

// 需要session对象

// MemorySession ...
type MemorySession struct {
	Sessionid string
	Data      map[string]interface{}
	rwlock    sync.RWMutex
}

// NewMemorySession ...
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		Sessionid: id,
		Data:      make(map[string]interface{}, 16),
	}
	return s
}

// Set ...
func (m *MemorySession) Set(key string, value interface{}) error {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.Data[key] = value
	return nil
}

// Get ...
func (m *MemorySession) Get(key string) (interface{}, error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	if value, ok := m.Data[key]; ok {
		return value, nil
	}
	err := fmt.Errorf("not find id:%s session", key)
	return nil, err
}

// Del ...
func (m *MemorySession) Del(key string) error {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.Data, key)
	return nil
}

// Save ...
func (m *MemorySession) Save() error {
	return nil
}
