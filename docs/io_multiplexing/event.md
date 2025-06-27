The Event interface acts as platform-agnostic abstraction that bridges the gap between different OS multiplexing mechanism

Data structure
- Fd: int (file descriptor identifier)
- Operation: uint32 (matches underlying system call event types)

Operations
- toNative: convert to platform-specific IO primitives
   + Linux: epoll
   + MacOS: kqueue

