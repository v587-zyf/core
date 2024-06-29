package http_server

type HttpOption struct {
	listenAddr string
}

type Option func(opts *HttpOption)

func NewHttpOption() *HttpOption {
	o := &HttpOption{}

	return o
}

func WithListenAddr(addr string) Option {
	return func(opts *HttpOption) {
		opts.listenAddr = addr
	}
}
