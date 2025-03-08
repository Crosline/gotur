package socket

import (
	"net"
	"syscall"
)

type BaseSocket struct {
	handle        int
	socketFamily  int
	socketType    int
	socketProto   int
	socketAddress syscall.SockaddrInet4
}

func newBaseSocket(family, socktype, proto int) (*BaseSocket, error) {
	handle, err := syscall.Socket(family, socktype, proto)
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
	Accept() (Socket, error)
	Close() error
	Receive([]byte) (int, error)
	Send([]byte) error
}

func (socket *BaseSocket) Close() error {
	return syscall.Close(socket.handle)
}

// ParseIPv4 converts a string IP address to [4]byte format
func ParseIPv4(ipStr string) ([4]byte, error) {
	ip := net.ParseIP(ipStr).To4()
	if ip == nil {
		return [4]byte{}, &net.AddrError{Err: "invalid IPv4 address", Addr: ipStr}
	}

	var addr [4]byte
	copy(addr[:], ip)
	return addr, nil
}
