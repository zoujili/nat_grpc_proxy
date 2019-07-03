package client

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"log"
	"nats_grpc_proxy"
	pb "nats_grpc_proxy/proto"
	"nats_grpc_proxy/queue"
)


type NATSQueue struct {
	consume chan pb.PushEvent
	produce chan  pb.RestEvent

	con  *nats.Conn
}

func (n *NATSQueue) ConsumeQueue() chan pb.PushEvent{
	return n.consume
}
func (n *NATSQueue) ProduceQueue() chan pb.RestEvent{
	return n.produce
}

//consume something from readQueue
func (n *NATSQueue) PopConsumeQueue() pb.PushEvent {
	m := <-n.consume
	log.Printf("client pop message from read queue: %s", m)
	err := n.con.Publish(nats_grpc_proxy.ClientNATSPubChannel, []byte(m.String()))
	if err != nil {
		fmt.Println(err)
	}
	return m
}

//push  something to writeQueue
func  (n *NATSQueue) PushProduceQueue(event  pb.RestEvent) {
	select {
	case n.produce <-event:
		log.Println("client push message to write queue")
		return
	default:
		return
	}
}



func NewClientNATSQueue() queue.ClientQueue{
	con, err := nats.Connect(nats_grpc_proxy.ClientNatDefaultURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	n:=&NATSQueue{
		produce: make(chan pb.RestEvent,10),
		consume: make(chan pb.PushEvent,10),
	}
	n.con = con

	_, err = con.Subscribe(nats_grpc_proxy.ClientNATSSubChannel, func(m *nats.Msg) {
		n.PushProduceQueue(pb.RestEvent{ClientID:ClientID,Message:string(m.Data)})
	})
    return  n
}

