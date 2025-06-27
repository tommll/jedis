package main

import (
	"fmt"
	"jedis/internal/server"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	fmt.Println("Starting Jedis server...")

	var signals = make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	var wg sync.WaitGroup
	wg.Add(2)

	go server.RunAsyncTCPServer(&wg)

	// listen for termination signals
	go server.WaitForSignal(&wg, signals)

	wg.Wait()
}
