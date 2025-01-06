package server

import (
	"log"
	"net"
	"time"
)

func main() {
	TCPListener()
}

func TCPListener() {
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatalf("Error listening", err)
	}

	defer func() {
		err = listener.Close()

		if err != nil {
			log.Fatalf("Eror closing listener %v", err)
		}
	}()

	log.Printf("Listening the server at %s", listener.Addr())

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf("Eror accpeting error %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	now := time.Now()
	defer func() {
		err := conn.Close()

		if err != nil {
			log.Fatalf("Error closing connection %v", err)
		}

		log.Printf("close connection from %s. Connection duration %v ms", conn.RemoteAddr(), time.Since(now))
	}()

	log.Printf("Accpeted the connection %s", conn.RemoteAddr())

	_, err := conn.Write([]byte("Hello, World from the client"))

	if err != nil {
		log.Fatalf("Error wrting error: %v", err)
	}
}
