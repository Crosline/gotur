package gotur

import (
	s "github.com/crosline/gotur/socket"
)

type BaseServer struct {
	socket    s.Socket
	isRunning bool
	handler   func(s.Socket)
}

type Server interface {
	Start(address string, port int) error
	Stop() error
	Handle(handler func(s.Socket))
	IsRunning() bool
}

func NewBaseServer(socket s.Socket) *BaseServer {
	return &BaseServer{
		socket:    socket,
		isRunning: false,
	}
}

func (server *BaseServer) Handle(handler func(s.Socket)) {
	server.handler = handler
}

func (server *BaseServer) Stop() error {
	server.isRunning = false
	return server.socket.Close()
}

func (server *BaseServer) IsRunning() bool {
	return server.isRunning
}