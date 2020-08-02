package session

// Session ...
type Session interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	Del(string) error
	Save() error
}
