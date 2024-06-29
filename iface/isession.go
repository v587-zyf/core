package iface

import (
	"context"
	"net"
)

type ISession interface {
	Set(key string, value any)
	Get(key string) (any, bool)
	Remove(key string)

	GetID() uint64
	SetID(id uint64)

	Start()
	Close() error

	GetConn() net.Conn
	GetCtx() context.Context

	Send(msgID uint16, tag uint32, userID uint64, msg IProtoMessage) error
}

type ISessionMgr interface {
	Length() int
	GetOne(UID uint64) ISession
	IsOnline(UID uint64) bool

	Add(ss ISession)
	Disconnect(SID uint64)

	Once(UID uint64, fn func(mgr ISession))
	Range(fn func(uint64, ISession))
}

type ISessionMethod interface {
	Start(ss ISession)
	Recv(conn ISession, data any)
	Stop(ss ISession)
}
