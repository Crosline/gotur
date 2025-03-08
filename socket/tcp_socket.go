package socket

import (
	"syscall"
)

type TCPSocket struct {
	BaseSocket
}

func NewTCPSocket() (*TCPSocket, error) {
	socket, err := newBaseSocket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		return nil, err
	}

	return &TCPSocket{BaseSocket: *socket}, nil
}

func (socket *TCPSocket) Bind(address string, port int) error {
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

func (socket *TCPSocket) Listen() error {
	return syscall.Listen(socket.handle, syscall.SOMAXCONN)
}

func (socket *TCPSocket) Accept() (Socket, error) {
	conn, _, err := syscall.Accept(socket.handle)
	if err != nil {
		return nil, err
	}

	clientSocket := &TCPSocket{
		BaseSocket: BaseSocket{
			handle:        conn,
			socketFamily:  socket.socketFamily,
			socketType:    socket.socketType,
			socketProto:   socket.socketProto,
			socketAddress: socket.socketAddress,
		},
	}

	return clientSocket, nil
}

func (socket *TCPSocket) Receive(buffer []byte) (int, error) {
	n, err := syscall.Read(socket.handle, buffer)
	return n, err
}

func (socket *TCPSocket) Send(data []byte) error {
	_, err := syscall.Write(socket.handle, data)
	return err
}

func (socket *TCPSocket) SetRemoteAddress(addr syscall.SockaddrInet4) {
	socket.socketAddress = addr
}
