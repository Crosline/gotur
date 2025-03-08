package socket

import (
	"errors"
)

type TCPServer struct{}

func (t *TCPServer) Bind(address string) error {
	return nil
}

func (t *TCPServer) Listen() error {
	return nil
}

func (t *TCPServer) Accept() (*TCPServer, error) {
	return nil, errors.New("not implemented")
}

func (t *TCPServer) Close() error {
	return nil
}

func (t *TCPServer) Read(b []byte) (int, error) {
	return 0, nil
}

func (t *TCPServer) Write(b []byte) (int, error) {
	return 0, nil
}
