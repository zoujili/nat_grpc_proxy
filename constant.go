package nats_grpc_proxy

//client nats to server nats
const ClientNATSSubChannel  =  "client.receive.upstream"
const ServerNATSPubChannel  =  "server.push.downstream"

//server nats to client nats
const ServerNATSSubChannel  =  "server.receive.upstream"
const ClientNATSPubChannel  =  "client.push.downstream"



const ClientNatDefaultURL              = "nats://127.0.0.1:4222"
const ServerNatDefaultURL              = "nats://127.0.0.1:4223"