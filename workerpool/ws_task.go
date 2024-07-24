package workerpool

import (
	"core/iface"
	"core/log"
	"core/session/ws_session"
	"go.uber.org/zap"
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
	if t.Func == nil {
		log.Warn("ws task func is nil", zap.Uint16("msgID", t.Data.(*iface.MessageFrame).MsgID))
		return
	}
	t.Func(t.Session, t.Data)
}
