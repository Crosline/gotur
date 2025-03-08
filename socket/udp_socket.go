package socket

import (
	"syscall"
)

type UDPSocket struct {
	BaseSocket
}

func NewUDPSocket() (*UDPSocket, error) {
	socket, err := newBaseSocket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		return nil, err
	}

	return &UDPSocket{BaseSocket: *socket}, nil
}

func (socket *UDPSocket) Bind(address string, port int) error {
	if socket.handle == 0 {
		return syscall.EINVAL
	}

	addr, err := ParseIPv4(address)
	if err != nil {
		return err
	}

	socketAddress := syscall.SockaddrInet4{Port: port, Addr: addr}
	socket.socketAddress = socketAddress

	return syscall.Bind(socket.handle, &socketAddress)
}

func (socket *UDPSocket) Listen() error {
	// UDP is connectionless, no need to listen
	return nil
}

func (socket *UDPSocket) Accept() (Socket, error) {
	// UDP is connectionless, just return self
	return socket, nil
}

func (socket *UDPSocket) Receive(buffer []byte) (int, error) {
	n, _, err := syscall.Recvfrom(socket.handle, buffer, 0)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (socket *UDPSocket) Send(data []byte) error {
	return syscall.Sendto(socket.handle, data, 0, &socket.socketAddress)
}

func (socket *UDPSocket) SetRemoteAddress(addr syscall.SockaddrInet4) {
	socket.socketAddress = addr
}
