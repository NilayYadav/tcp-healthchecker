package main

import (
	"time"

	"github.com/nilay/tcp-server/checker"
	"github.com/nilay/tcp-server/server"
)

func main() {
	go server.TCPListener()

	time.Sleep(1 * time.Second)

	checker.RunChecker()
}
