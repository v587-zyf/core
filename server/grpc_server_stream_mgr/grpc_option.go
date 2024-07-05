package grpc_server_stream_mgr

type GrpcOption struct {
	id uint64
}

type Option func(opts *GrpcOption)

func NewGrpcOption() *GrpcOption {
	o := &GrpcOption{}

	return o
}

func WithID(id uint64) Option {
	return func(opts *GrpcOption) {
		opts.id = id
	}
}
