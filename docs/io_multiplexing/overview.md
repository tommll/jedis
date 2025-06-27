### Traditional multi-threaded approach
- Thread 1 -> Client 1 socket
- Thread 2 -> Client 2 socket
...
- Thread N -> Client N socket

### I/O Multiplexed approach
- Single thread -> I/O Multiplexer -> Multiple sockets
