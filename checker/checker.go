package checker

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func RunChecker() {
	checker := NewTCPChecker(net.ParseIP("127.0.0.1"), 3000, 10)
	checker.Timeout = 1 * time.Second

	logOutput := log.Writer()

	result := checker.CheckRetries(120, 2*time.Second, logOutput)

	println("Health Check Result: ", result.Message)
}

type Target struct {
	IP      net.IP
	Port    int
	Packets int
}

type TCPChecker struct {
	Target
	Timeout time.Duration
}

type Result struct {
	Success bool
	Message string
}

func NewTCPChecker(ip net.IP, port int, packets int) *TCPChecker {
	return &TCPChecker{
		Target: Target{
			IP:      ip,
			Port:    port,
			Packets: packets,
		},
	}
}
func (hc *TCPChecker) addr() string {
	return fmt.Sprintf("%s:%d", hc.IP.String(), hc.Port)
}

func (hc *TCPChecker) Check(timeout time.Duration) *Result {

	conn, err := net.DialTimeout("tcp", hc.addr(), timeout)

	println("Connected to ", hc.addr())

	if err != nil {
		return &Result{Success: false, Message: fmt.Sprintf("Failed to connect to %s", err)}
	}

	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by server")
				break
			}

		}

		log.Printf("Received %q", buf[:n])
	}
	return &Result{Success: true, Message: fmt.Sprintf("Connected to %s", hc.addr())}
}

func (hc *TCPChecker) CheckRetries(retries int, retryDelay time.Duration, logOutput io.Writer) *Result {
	var result *Result

	for i := 0; i < retries; i++ {
		start := time.Now()
		result = hc.Check(hc.Timeout)
		duration := time.Since(start)

		logOutput.Write([]byte(fmt.Sprintf("Health Check %d - Success: %v, Latency: %v, MSG: %s\n", i+1, result.Success, duration, result.Message)))

		if result.Success {
			return result
		}

		time.Sleep(retryDelay)
	}

	return result
}
