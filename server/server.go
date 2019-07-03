package server

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	pb "nats_grpc_proxy/proto"
	"nats_grpc_proxy/queue"
	"net"
)

type Server struct {
	Queue queue.ServiceQueue
}

func newServer() *Server {
	return &Server{}
}

func (s *Server) Say(stream pb.Chat_SayServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("client send end")
			return nil
		}
		if err != nil {
			log.Println("receive client error:", err)
			return err
		}
		Pool.Add(resp.ClientID,stream)
		go func() {
			<-stream.Context().Done()
			Pool.Del(resp.ClientID)
		}()

		if resp != nil {
			fmt.Println("receive message form client",resp)
			s.Queue.ConsumeQueue() <- *resp
		}
	}
}

func (s *Server) Push() {
	for {
		m := <-s.Queue.ProduceQueue()
		log.Printf("push message to  client: %s", m.GetMessage())
		Pool.Range(func(clientID, stream interface{}) bool {
			id :=clientID.(string)
			st :=stream.(pb.Chat_SayServer)
			if m.GetClientID() == id {
				if err := st.Send(&m); err != nil {
					log.Printf("send failed: %s", err)
				}
				return true
			}
			return true
		})
	}
}

func (c *Server) Consume(){
	for {
		m := c.Queue.PopConsumeQueue()
		log.Printf("Success Receive Process: %s", m)
	}
}

func NewServer() {
	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	s := grpc.NewServer()
	q,err := NewNATSQueue()
	if err !=nil {
		log.Fatalf("failed to connect with nats: %s", err)
	}
	my := newServer()
	my.Queue = q

	pb.RegisterChatServer(s, my)

	go my.Push()
	go my.Consume()

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}