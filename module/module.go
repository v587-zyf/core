package module

import (
	"context"
	"core/iface"
)

type DefModule struct{}

func (m *DefModule) Name() string {
	return ""
}

func (m *DefModule) Init(ctx context.Context, opts ...iface.Option) error {
	return nil
}

func (m *DefModule) Start() error {
	return nil
}

func (m *DefModule) Run() {}

func (m *DefModule) Stop() {}
