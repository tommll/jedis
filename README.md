
## How to run
```
  cd cmd
  go run main.go
  # on another terminal
  redis-cli -p 8080
```

#### Supported features
- Compatible with Redis CLI
- Single-threaded architecture
- Multiplexing IO using epoll for Linux and kqueue for MacOS
- RESP protocol
- Graceful shutdown
- Simple eviction mechanism


#### Commands
- PING
- SET, GET, TTL

#### Data types
- String

#### Data structures
- Hash table

## Todos
- Other data types: number, array, boolean, date, ...
- Other eviction strategies: LRU, LFU
