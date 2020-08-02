package session

// SessionMgr ...
type SessionMgr interface {
	Init(string, ...string)
	CreateSession() (Session, error)
	Get(string) (Session, error)
}
