package io_multiplexing

type Operation int

const (
	OpRead  = 0
	OpWrite = 1
)

type Event struct {
	Fd int
	Op Operation
}

type IOMultiplexer interface {
	Monitor(event Event) error
	CheckReadyEvents() ([]Event, error)
	Close() error
}
