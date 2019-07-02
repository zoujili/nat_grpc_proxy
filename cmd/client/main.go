package main

import "nats_grpc_proxy/client"
func main() {
	go client.NewClient()
	select {
	}
}

