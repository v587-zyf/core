package servermanager

import (
	"cwcommon/protobuf/pbregister"
	"cwcore/cnet"
	"cwcore/errcode"
	"cwcore/log"
	"google.golang.org/grpc"
	"sync"

	"go.uber.org/zap"
)

type ServerInfo struct {
	SID        int32
	ServerType string
	LinkAddr   string
	Ver        string
	Server     *grpc.ClientConn
}

type Fn func(sid int32, serverInfo *ServerInfo) bool

func NewServerManager() *ServerManager {
	return &ServerManager{}
}

type ServerManager struct {
	servers sync.Map
}

func (mgr *ServerManager) Registr(ss cnet.ISession, info *pbregister.ServerInfo) bool {
	_, loaded := mgr.servers.LoadOrStore(*info.Sid, &ServerInfo{
		SID:        *info.Sid,
		ServerType: *info.Type,
		LinkAddr:   *info.Addr,
		Ver:        *info.Ver,
		Session:    ss,
	})

	if loaded {
		log.Error("server has existed", zap.Int32p("serverID", info.Sid))
		return false
	}

	// succeed
	ss.SetID(*info.Sid)

	servers := make([]*pbregister.ServerInfo, 0)
	// notify other server
	mgr.servers.Range(func(key, value any) bool {
		serverInfo := value.(*ServerInfo)
		if serverInfo.SID == *info.Sid { // not notify self
			return true
		}
		ntf := new(pbregister.RegisterUpdateNtf)
		ntf.Info = info
		serverInfo.Session.SendMsg(int32(pbregister.MSGID_REGISTER_UPDATE_NTF), 0, ss.GetID(), ntf)

		servers = append(servers, &pbregister.ServerInfo{
			Sid:  &serverInfo.SID,
			Type: &serverInfo.ServerType,
			Ver:  &serverInfo.Ver,
			Addr: &serverInfo.LinkAddr,
		})

		return true
	})

	if len(servers) > 0 {
		ntf := new(pbregister.RegisterTotalNtf)
		ntf.Infos = servers
		ss.SendMsg(int32(pbregister.MSGID_REGISTER_TOTAL_NTF), 0, ss.GetID(), ntf)
	}

	log.Info("register server succ", zap.Int32("serverID", ss.GetID()), zap.String("serverType", *info.Type), zap.String("linkAddr", *info.Addr), zap.String("ver", *info.Ver))

	return true
}

func (mgr *ServerManager) Unregister(ss cnet.ISession) bool {
	v, ok := mgr.servers.LoadAndDelete(ss.GetID())
	if !ok {
		log.Error("server not existed", zap.Int32("serverID", ss.GetID()))
		return false
	}

	info := v.(*ServerInfo)

	ack := new(pbregister.RegisterDeleteNtf)

	ack.Info = &pbregister.ServerInfo{
		Sid:  &info.SID,
		Type: &info.ServerType,
		Ver:  &info.Ver,
		Addr: &info.LinkAddr,
	}
	mgr.Range(func(sid int32, serverInfo *ServerInfo) bool {
		err := serverInfo.Session.SendMsg(int32(pbregister.MSGID_REGISTER_DELETE_NTF), 0, 0, ack)
		if err != nil {
			log.Error("send register delete ntf failed", zap.Int32("serverID", sid), zap.String("err", err.Error()))
		}
		return true
	})

	return ok
}

func (mgr *ServerManager) Range(fn Fn) {
	mgr.servers.Range(func(key, value any) bool {
		sid := key.(int32)
		serverInfo := value.(*ServerInfo)
		return fn(sid, serverInfo)
	})
}

func (mgr *ServerManager) One(sid int32, fn Fn) error {
	v, ok := mgr.servers.Load(sid)
	if !ok {
		return errcode.ERR_SERVER_NOT_FOUND
	}
	fn(sid, v.(*ServerInfo))

	return nil
}
