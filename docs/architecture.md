
# Architecture
### 1. Single-threaded event loop
The system follows a single-threaded, event-driven architecture

- Main Event Loop: handle all client connections and commands in a single thread
- I/O multiplexing: use X and Y for efficient event handling
- Non-blocking I/O: all file descriptors operate in non-blocking mode
- Atomic state management: uses atomic operations for thread-safe state transitions

### 2. Network layer

- TCP server
  + Asynchronous TCP server: handle multiple client connection concurrently
  + Connection management: tracks active connections and handles client lifecycle
  + Graceful shutdown: implements proper cleanup on termination signals

### 3. Protocol layer

#### RESP protocol implementation
- Redis compatibility
- Command parsing
- Response encoding
- Support types:
  + Simple strings
  + Integers
  + Bulk strings
  + Arrays
  + Errors
  + Custom Integer Arrays

#### Command structure
```go
type KVCmd struct {
    Cmd  string
    Args []string
}
```

### 4. Storage Engine

#### Dictionay implementation
- Primary storage
- Object System
- TTL Support
- Eviction Strategy
- Memory Management

#### Object Types and Encoding
```go
type Obj struct {
    Value        interface{} // Actual data
    TypeEncoding uint8       // Type and encoding flags
}
```


### 5. Configuration system

#### Config management
- Runtime config (port, host, connection limits)
- Storage limits: maximum number of keys
- Eviction strategy: configurable eviction policies
- Persistence: AOF (Append-only file) configuration


### 6. State Management

#### Engine status
- Status tracking: atomic state management for graceful shutdown
- States:
  + Waiting
  + Busy
  + ShuttingDown

# Performance characteristics

### Concurrency model
- Single threaded: remove lock contention and context switching overhead
- Event-driven: efficient handling of multiple concurrent connections
- Non-blocking I/O: threads not blocking on I/O operations

### Memory management
- In-memory storage: all data stored in RAM
- Automatic expiration: TTL-based cleanup of expired keys
- Eviction Policies: memory pressure handling through configurable eviction

### Scalability
- Connection Limits: configurable maximum connections (default: 20,000)
- Key Limits: configurable maximum keys (default: 5,000,000)

# Design Patten

### 1. Event-driven architecutre
- Reactor patten: single thread handles multiple I/O sources
- Non-blocking operations: all I/O operations are non-blocking
- Event loop: continuous event processing loop

### 2. Strategy patten
- Eviction strategy: pluggable eviction policies
- I/O multiplexing multi-platform support: platform-specific implementations behind common interface

### 3. Command pattern
- Command dispatch: centralized command routing
- Handler separation: each command type has its own handlers




