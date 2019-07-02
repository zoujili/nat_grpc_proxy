package server

import (
	"log"
	pb "nats_grpc_proxy/proto"
	"nats_grpc_proxy/queue"
)

type NATSQueue struct {
	consume chan pb.RestEvent
	produce chan pb.PushEvent
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
	n:=&NATSQueue{
		consume: make(chan pb.RestEvent,10),
		produce: make(chan pb.PushEvent,10),
	}
    return  n
}

