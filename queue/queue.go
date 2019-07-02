package queue

import (
	pb "nats_grpc_proxy/proto"
)

type ServiceQueue interface {
	ConsumeQueue() chan pb.RestEvent
	ProduceQueue()  chan pb.PushEvent

	PopConsumeQueue()  pb.RestEvent
	PushProduceQueue(event  pb.PushEvent)
}

type ClientQueue interface {
	ConsumeQueue() chan pb.PushEvent
	ProduceQueue()  chan pb.RestEvent

	PopConsumeQueue()  pb.PushEvent
	PushProduceQueue(event  pb.RestEvent)
}