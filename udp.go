package gotur

import (
	"errors"
	"strings"
	"syscall"
)

type UDPServer struct {
	socket syscall.Handle
}

func (u *UDPServer) Bind(ip string, port int) error {
	handle, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		return err
	}

	u.socket = handle

	ipBytes := strings.Split(ip, ".")
	if len(ipBytes) != 4 {
		return errors.New("invalid ip address")
	}

	ipAddr := [4]byte{}
	for i, b := range ipBytes {
		ipAddr[i] = byte.parse(b)
	}

	addr := syscall.SockaddrInet4{Port: port, Addr: ip}
	return syscall.Bind(handle, &addr)
}

func (u *UDPServer) Listen() error {
	// UDP is connectionless, so listen simply means waiting for packets
	return nil
}

func (u *UDPServer) Accept() (*UDPServer, error) {
	// Not applicable for UDP
	return nil, errors.New("UDP does not support Accept")
}

func (u *UDPServer) Close() error {
	return syscall.Close(u.socket)
}

func (u *UDPServer) Read(b []byte) (int, *syscall.SockaddrInet4, error) {
	n, from, err := syscall.Recvfrom(u.socket, b, 0)
	if err != nil {
		return 0, nil, err
	}

	addr, ok := from.(*syscall.SockaddrInet4)
	if !ok {
		return 0, nil, errors.New("invalid address type")
	}

	return n, addr, nil
}

func (u *UDPServer) Write(b []byte, addr *syscall.SockaddrInet4) error {
	return syscall.Sendto(u.socket, b, 0, addr)
}