package socket

import (
	"errors"
	"strings"
	"syscall"
)

type UDPSocket struct {
	BaseSocket
}

func NewUDPSocket() (*UDPSocket, error) {
	socket, err := NewBaseSocket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		return nil, err
	}

	return &UDPSocket{BaseSocket: *socket}, nil
}

func (socket *UDPSocket) Bind(address string, port int) error {
	if socket.handle == 0 {
		return errors.New("socket is not initialized")
	}
	
	ipBytes := strings.Split(address, ".")
	if len(ipBytes) != 4 {
		return errors.New("invalid ip address")
	}

	addr := [4]byte{}
	for i, b := range ipBytes {
		addr[i] = []byte(b)[0]
	}

	socketAddress := syscall.SockaddrInet4{Port: port, Addr: addr}
	socket.socketAddress = socketAddress

	return syscall.Bind(socket.handle, &socketAddress)
}

func (socket *UDPSocket) Listen() error {
	return nil
}

func (socket *UDPSocket) Accept() (*UDPSocket, error) {
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