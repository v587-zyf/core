package workerpool

import (
	"core/iface"
	"core/session/ws_session"
)

func (p *WorkerPool) AssignWsTask(fn ws_session.Recv, ss iface.IWsSession, data any) error {
	return Assign(&WsTask{
		Func:    fn,
		Session: ss,
		Data:    data,
	})
}

type WsTask struct {
	Func    ws_session.Recv
	Session iface.IWsSession
	Data    any
}

func (t *WsTask) Do() {
	t.Func(t.Session, t.Data)
}
