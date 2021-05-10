package sync

type CommandType uint8

const (
	STOP CommandType = iota
	FAIL
)

type Command struct {
	Command CommandType
	Message interface{}
}
