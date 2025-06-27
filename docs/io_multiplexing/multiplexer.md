### Why need to monitor?
The ultimate goal is to receive data from a "client". While we can wait until the client returns some data, it inefficient.

Instead, we can add let another service do the waiting for us. When the data is ready, that service will notify us.

### Core operations
- Register an event for monitoring
- Find ready events (events that have data available)
- Stop all monitoring operations

### Core interface
```go
type IOMultiplexer interface {
	Monitor(event Event) error
	CheckReadyEvents() ([]Event, error)
	Close() error
}
```

### Usage
- At a regular interval, check for ready events
- For each ready event, process them
