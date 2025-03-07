package socket

import "syscall"

type BaseSocket struct {
	handle            syscall.Handle
	socketFamily  int
	socketType    int
	socketProto   int
	socketAddress syscall.SockaddrInet4
}

func newBaseSocket(family, socktype, proto int) (*BaseSocket, error) {
	handle, err := syscall.Socket(family, socktype, syscall.IPPROTO_UDP)
	if err != nil {
		return nil, err
	}
	
	return &BaseSocket{
		handle:       handle,
		socketFamily: family,
		socketType:   socktype,
		socketProto:  proto,
	}, nil
}

type Socket interface {
	Bind(string, int) error
	Listen() error
	Accept() (*BaseSocket, error)
	Close() error
	Receive([]byte) (int, *syscall.SockaddrInet4, error)
	Send([]byte) error
}


func (socket *BaseSocket) Close() error {
	return syscall.Close(socket.handle)
}