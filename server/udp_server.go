package gotur

import (
	s "github.com/crosline/gotur/socket"
)

type UDPServer struct {
	BaseServer
}

func NewUDPServer() (*UDPServer, error) {
	socket, err := s.NewUDPSocket()
	if err != nil {
		return nil, err
	}
	
	baseServer := NewBaseServer(socket)
	return &UDPServer{
		BaseServer: *baseServer,
	}, nil
}

func (server *UDPServer) Start(address string, port int) error {
	if err := server.socket.Bind(address, port); err != nil {
		return err
	}
	
	server.isRunning = true
	
	go func() {
		buffer := make([]byte, 4096)
		for server.isRunning {
			n, err := server.socket.Receive(buffer)
			if err != nil {
				// Could log error here
				continue
			}
			
			if server.handler != nil {
				// Create a copy of the buffer to avoid data races
				data := make([]byte, n)
				copy(data, buffer[:n])
				
				go server.handler(server.socket)
			}
		}
	}()
	
	return nil
}