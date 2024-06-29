package server

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type HttpServer struct {
	addr string
	pem  string
	key  string

	upGrader *websocket.Upgrader
}

func NewServer(addr, pem, key string) *HttpServer {
	s := &HttpServer{
		addr: addr,
		pem:  pem,
		key:  key,
		upGrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 4096,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	return s
}

func (s *HttpServer) Start() {
	http.HandleFunc("/ws", s.wsHandle)
	err := http.ListenAndServeTLS(s.addr, s.pem, s.key, nil)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) wsHandle(w http.ResponseWriter, r *http.Request) {
	// 设置Access-Control-Allow-Origin头部，允许跨域请求
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// 处理跨域请求
	if r.Method == "OPTIONS" {
		// 处理预检请求
		w.WriteHeader(http.StatusOK)
		return
	}
	// websocket
	// 1.http升级为websocket
	//wsConn, err := s.upGrader.Upgrade(w, r, nil)
	//if err != nil {
	//	log.Error("webSocket upgrade err:", zap.Error(err))
	//}

	// todo 这里有递归引用 需要解决
	// 2.连接后 s或c端都可收发消息
	//ss := session.NewSession(context.Background(), wsConn)
	//global.SessionMgr.Add(ss)
	//ss.Start()
}
