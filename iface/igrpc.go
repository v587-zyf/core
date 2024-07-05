package iface

import (
	"context"
	"google.golang.org/grpc"
)

type IGrpc interface {
	Init(ctx context.Context, opts ...any) error
	Start()
}

type IGrpcStream interface {
	Init(ctx context.Context, opts ...any) error
	SetStream(stream *grpc.ClientStream)

	GetID() uint64
	GetStream() *grpc.ClientStream
	GetCtx() context.Context
}
