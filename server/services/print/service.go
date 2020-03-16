package print

import (
	"bufio"
	"fmt"
	"net"

	"github.com/edwintcloud/print-solution/server/repositories"
)

// Service represents a print service
type Service struct {
	clients repositories.Repository
	jobs    repositories.Repository
}

// NewPrintService creates a new print service
func NewPrintService(clientsRepository, jobsRepository repositories.Repository) *Service {
	return &Service{
		clients: clientsRepository,
		jobs:    jobsRepository,
	}
}

// PingClient performs tcp ACK on specified client address
func (p *Service) PingClient(clientAddr string) {
	conn, _ := net.Dial("tcp", clientAddr)
	fmt.Fprintf(conn, "ACK\n")
	resp, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Response from server: %s\n", resp)

	// Note: we need to use DialWithTimeout as well as creating a context with timeout for the tcp ack,
	// we can ping all services concurrently using go routines and a waitgroup (ref: https://developer20.com/tcp-scanner-in-go/)
}
