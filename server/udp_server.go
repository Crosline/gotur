package gotur

import (
	s "github.com/crosline/gotur/socket"
)


func NewUDPServer() (*BaseServer, error) {
	// Create the UDP socket internally
	socket, err := s.NewUDPSocket()
	if err != nil {
		return nil, err
	}
	
	// Note: socket implements the Socket interface, so this works
	baseServer := NewBaseServer(socket)
}

// Start starts the UDP server
func (s *UDPServer) Start(address string, port int) error {
	if err := s.socket.Bind(address, port); err != nil {
		return err
	}
	
	s.isRunning = true
	
	go func() {
		buffer := make([]byte, 4096)
		for s.isRunning {
			n, err := s.socket.Receive(buffer)
			if err != nil {
				// Log error or handle it
				continue
			}
			
			if s.handler != nil {
				// Create a copy of the buffer to avoid data races
				data := make([]byte, n)
				copy(data, buffer[:n])
				
				go func() {
					// For UDP, we just pass the original socket
					s.handler(s.socket)
				}()
			}
		}
	}()
	
	return nil
}