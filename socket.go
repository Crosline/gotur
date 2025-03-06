package gotur

import "syscall"

type Socket interface {
	Bind(ip string, port int) error
	Listen() error
	Accept() (Socket, error)
	Close() error
	Read([]byte) (int, *syscall.SockaddrInet4, error)
	Write([]byte) error
}