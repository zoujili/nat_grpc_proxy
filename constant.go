package nats_grpc_proxy

//client nats to server nats
const ClientNATSSubChannel  =  "delaware.sub"
const ServerNATSPubChannel  =  "fitstation.pub"

//server nats to client nats
const ServerNATSSubChannel  =  "fitstation.sub"
const ClientNATSPubChannel  =  "delaware.pub"



const ClientNatDefaultURL              = "nats://127.0.0.1:4222"
const ServerNatDefaultURL              = "nats://127.0.0.1:4223"