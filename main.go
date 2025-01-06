package main

import (
	"NilayYadav/tcp-healthchecker/checker"
	"NilayYadav/tcp-healthchecker/server"
	"time"
)

func main() {
	go server.TCPListener()
	time.Sleep(1 * time.Second)
	checker.RunChecker()
}
