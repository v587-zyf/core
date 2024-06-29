package workerpool

import (
	"core/iface"
	"core/session"
	"sync"
)

type ITask interface {
	Do()
}

var defaultWorkPoll *WorkerPool
var once sync.Once

func Init(cfg ...*Config) (err error) {
	once.Do(func() {
		defaultWorkPoll, err = New(cfg...)
		defaultWorkPoll.Start()
	})

	return
}

func Assign(task ITask) error {
	return defaultWorkPoll.Assign(task)
}

func AssignNetTask(fn session.Recv, ss iface.ISession, data any) error {
	return defaultWorkPoll.Assign(&NetTask{
		Func:    fn,
		Session: ss,
		Data:    data,
	})
}
