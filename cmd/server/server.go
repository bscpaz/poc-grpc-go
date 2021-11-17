//This is just a server and UserServer must be registred in this sevrer.

package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	listenerToPort, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(listenerToPort); err != nil {
		log.Fatalf("Could not serve %v", err)
	}
}
