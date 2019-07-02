package main

import "nats_grpc_proxy/server"
func main() {
	go server.NewServer()
	select {
	}
}

