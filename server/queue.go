package server

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"log"
	"nats_grpc_proxy"
	pb "nats_grpc_proxy/proto"
	"nats_grpc_proxy/queue"
)

type NATSQueue struct {
	consume chan pb.RestEvent
	produce chan pb.PushEvent

	con  *nats.Conn
}

func (n *NATSQueue) ConsumeQueue() chan pb.RestEvent{
	return n.consume
}
func (n *NATSQueue) ProduceQueue() chan pb.PushEvent{
	return n.produce
}

//consume something from readQueue
func (n *NATSQueue) PopConsumeQueue() pb.RestEvent{
	m := <-n.consume
	log.Printf("server pop message from consume queue: %s", m)
	err := n.con.Publish(nats_grpc_proxy.ServerNATSPubChannel, []byte(m.String()))
	if err != nil {
		fmt.Println(err)
	}
	return  m
}

//push  something to writeQueue
func  (n *NATSQueue) PushProduceQueue(event  pb.PushEvent) {
	select {
	case n.produce <-event:
		log.Println("server push message to produce queue")
		return
	default:
		return
	}
}



func NewNATSQueue() queue.ServiceQueue{
	con, err := nats.Connect(nats_grpc_proxy.ServerNatDefaultURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	n:=&NATSQueue{
		consume: make(chan pb.RestEvent,10),
		produce: make(chan pb.PushEvent,10),
	}
	n.con = con

	//todo  with client id
	_, err = con.Subscribe(nats_grpc_proxy.ServerNATSSubChannel, func(m *nats.Msg) {
		n.PushProduceQueue(pb.PushEvent{ClientID:"test",Message:string(m.Data)})
	})
    return  n
}

