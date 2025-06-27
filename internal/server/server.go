package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"syscall"

	"jedis/internal/config"
	"jedis/internal/constant"
	"jedis/internal/core"
	"jedis/internal/core/io_multiplexing"
)

func RunAsyncTCPServer(wg *sync.WaitGroup) error {
	defer wg.Done()

	fmt.Printf("Server started on port %d\n", config.Port)
	var events = make([]io_multiplexing.Event, config.MaxConnections)
	clientNumber := 0

	// Create a server socket. A socket is a way to communicate between client and server.
	serverFD, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("Error creating server socket:", err)
		return err
	}

	defer syscall.Close(serverFD)

	// Set the socket to operate in a non-blocking mode
	if err = syscall.SetNonblock(serverFD, true); err != nil {
		fmt.Println("Error setting non-blocking mode:", err)
		return err
	}

	// Bind the socket to a specific port
	ip4 := net.ParseIP(config.Host)
	if err = syscall.Bind(serverFD, &syscall.SockaddrInet4{
		Port: config.Port,
		Addr: [4]byte{ip4[0], ip4[1], ip4[2], ip4[3]},
	}); err != nil {
		fmt.Println("Error binding socket:", err)
		return err
	}

	// Listen for incoming connections
	if err = syscall.Listen(serverFD, config.MaxConnections); err != nil {
		fmt.Println("Error listening for connections:", err)
		return err
	}

	ioMultiplexer, err := io_multiplexing.CreateIOMultiplexer()
	if err != nil {
		fmt.Println("Error creating I/O multiplexer:", err)
		return err
	}
	defer ioMultiplexer.Close()

	// Monitor "read" events on the Server FD
	if err = ioMultiplexer.Monitor(io_multiplexing.Event{
		Fd: serverFD,
		Op: io_multiplexing.OpRead,
	}); err != nil {
		return err
	}

	// Check server status atomically
	for atomic.LoadInt32(&status) != constant.StatusShuttingDown {
		// check if any FD is ready for an IO
		events, err = ioMultiplexer.Check()
		if err != nil {
			continue
		}

		if !atomic.CompareAndSwapInt32(&status, constant.StatusWaiting, constant.StatusBusy) {
			if status == constant.StatusShuttingDown {
				return nil
			}
		}

		for i := 0; i < len(events); i++ {
			if events[i].Fd == serverFD {
				fmt.Println("New client connected")

				// the Server FD is ready for reading, means we have a new client.
				clientNumber++
				log.Printf("new client: id=%d\n", clientNumber)

				connFD, _, err := syscall.Accept(serverFD)
				if err != nil {
					log.Println("Error accepting connection:", err)
					continue
				}

				if err = syscall.SetNonblock(connFD, true); err != nil {
					log.Println("Error setting non-blocking mode:", err)
					return err
				}

				// Add new client connection to be monitored
				if err = ioMultiplexer.Monitor(io_multiplexing.Event{
					Fd: connFD,
					Op: io_multiplexing.OpRead,
				}); err != nil {
					log.Println("Error monitoring client connection:", err)
					return err
				}
			} else {
				// the Client FD is ready for reading, means an existing client is sending a command
				comm := core.FDComm{Fd: int(events[i].Fd)}
				cmd, err := readCommandFD(comm.Fd)

				if err != nil {
					syscall.Close(events[i].Fd)
					clientNumber--
					log.Println("client quit")
					atomic.SwapInt32(&status, constant.StatusWaiting)
					continue
				}
				responseRw(cmd, comm)
			}
			atomic.SwapInt32(&status, constant.StatusWaiting)
		}
	}

	return nil
}

var status int32 = constant.StatusWaiting

// Read from a file descriptor for command data
func readCommandFD(fd int) (*core.Cmd, error) {
	var buf = make([]byte, 512)
	n, err := syscall.Read(fd, buf)
	if err != nil {
		return nil, err
	}
	return core.ParseCmd(buf[:n])
}

func responseErrorRw(err error, rw io.ReadWriter) {
	rw.Write([]byte(fmt.Sprintf("-%s%s", err, constant.CRLF)))
}

func responseRw(cmd *core.Cmd, rw io.ReadWriter) {
	// Core command execution logic
	err := core.EvalAndResponse(cmd, rw)
	if err != nil {
		responseErrorRw(err, rw)
	}
}

func WaitForSignal(wg *sync.WaitGroup, signals chan os.Signal) {
	defer wg.Done()

	<-signals

	// ensure no commands are being processed
	for atomic.LoadInt32(&status) == constant.StatusBusy {
	}
	fmt.Println("Received termination signal, shutting down...")
	os.Exit(0)
}
