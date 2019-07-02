package client

import (
	"log"
	pb "nats_grpc_proxy/proto"
	"nats_grpc_proxy/queue"
)


type NATSQueue struct {
	consume chan pb.PushEvent
	produce chan  pb.RestEvent
}

func (n *NATSQueue) ConsumeQueue() chan pb.PushEvent{
	return n.consume
}
func (n *NATSQueue) ProduceQueue() chan pb.RestEvent{
	return n.produce
}

//consume something from readQueue
func (n *NATSQueue) PopConsumeQueue() pb.PushEvent{
	m := <-n.consume
	log.Printf("client pop message from read queue: %s", m)
	return  m
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
	n:=&NATSQueue{
		produce: make(chan pb.RestEvent,10),
		consume: make(chan pb.PushEvent,10),
	}
    return  n
}

