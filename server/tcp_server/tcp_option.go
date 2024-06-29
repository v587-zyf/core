package tcp_server

import (
	"core/iface"
)

type TcpOption struct {
	sid int64

	listenAddr string

	method iface.ISessionMethod
}

type Option func(opts *TcpOption)

func NewTcpOption() *TcpOption {
	o := &TcpOption{}

	return o
}

func WithListenAddr(addr string) Option {
	return func(opts *TcpOption) {
		opts.listenAddr = addr
	}
}

func WithSid(sid int64) Option {
	return func(opts *TcpOption) {
		opts.sid = sid
	}
}

func WithMethod(m iface.ISessionMethod) Option {
	return func(opts *TcpOption) {
		opts.method = m
	}
}
