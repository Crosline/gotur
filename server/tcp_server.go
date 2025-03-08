package gotur

import (
	s "github.com/crosline/gotur/socket"
)

type TCPServer struct {
	BaseServer
}

func NewTCPServer() (*TCPServer, error) {
	socket, err := s.NewTCPSocket()
	if err != nil {
		return nil, err
	}

	baseServer := NewBaseServer(socket)
	return &TCPServer{
		BaseServer: *baseServer,
	}, nil
}

func (server *TCPServer) Start(address string, port int) error {
	if err := server.socket.Bind(address, port); err != nil {
		return err
	}

	if err := server.socket.Listen(); err != nil {
		return err
	}

	server.isRunning = true

	go func() {
		for server.isRunning {
			clientSocket, err := server.socket.Accept()
			if err != nil {
				// Could log error here
				continue
			}

			if server.handler != nil {
				go server.handler(clientSocket)
			}
		}
	}()

	return nil
}
