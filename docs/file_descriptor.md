### What is a file descriptor?
- In Unix systems, file descriptors are integers that represent open files, pipes, sockets or other I/O resources. They provide a unified interface for all I/O operations.
- Unified I/O interface:
  + read() and write() for data transfer
  + close() for clean up
  + select(), poll() or epoll() for I/O multiplexing

### How is a file descriptor being used?

- When you create a socket with `syscall.Socket`, the OS
  + Allocates a new file descriptor
  + Associates it with the socket's internal data structures
  + Return the file descriptor number


