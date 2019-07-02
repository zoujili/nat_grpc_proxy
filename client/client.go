package client

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "nats_grpc_proxy/proto"
	"nats_grpc_proxy/queue"
	"time"
)

type Client struct {
	Queue queue.ClientQueue
}

func newClient() *Client {
	return &Client{}
}


func (c *Client) RegisterAndSay(stream pb.Chat_SayClient, name string) error {
	if err := stream.Send(&pb.RestEvent{Message: "ping", ClientID: name}); err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("server send end")
			return nil
		}

		if err != nil {
			log.Println("receive server error:", err)
			return err
		}

		if resp != nil {
			fmt.Println("receive message form server",resp)
			c.Queue.ConsumeQueue() <- *resp
		}

	}
	return nil
}

func (c *Client )Push(stream pb.Chat_SayClient) {
	for {
		select {
		case m := <-c.Queue.ProduceQueue():
			log.Printf("send message to server %s",m)
			if err := stream.Send(&m); err != nil {
				log.Println("err:", err)
			}
		}
	}

}

func NewClient() {
	addr := flag.String("a", "127.0.0.1:9092", "grpc server address")
	flag.Parse()
	// init important structures
	//rand.Seed(time.Now().UTC().UnixNano())
	//name := fmt.Sprintf("%d", rand.Intn(50))
	clientID := "test"

	q := NewClientNATSQueue()
	my := newClient()
	my.Queue = q


	// Setup a connection with the server
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewChatClient(conn)

	ctx := context.Background()
	stream, err := client.Say(ctx)
	if err != nil {
		log.Printf("create steam to server faile %s", err)
	}
	go my.RegisterAndSay(stream, clientID)
	go my.Push(stream)

	go func() {
		time.Sleep(6 * time.Second)
		my.Queue.PushProduceQueue(pb.RestEvent{ClientID:"hello",Message:"tom"})
	}()

	go func() {
		for {
			m := my.Queue.PopConsumeQueue()
			log.Printf("Success Receive Process: %s", m)
		}

	}()
}




