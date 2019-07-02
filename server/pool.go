package server

import "sync"

import  pb "nats_grpc_proxy/proto"

var Pool *ConnectPool

func init() {
	Pool = &ConnectPool{}
}

type ConnectPool struct {
	sync.Map
}

func (p *ConnectPool) Get(clientID string) pb.Chat_SayServer {
	if stream, ok := p.Load(clientID); ok {
		return stream.(pb.Chat_SayServer)
	} else {
		return nil
	}
}

func (p *ConnectPool) Add(clientID string, stream pb.Chat_SayServer) {
	p.Store(clientID, stream)
}


func (p *ConnectPool) Del(clientID string) {
	p.Delete(clientID)
}



