//This is just a server and UserServer must be registred in this sevrer.

package main

import (
	"log"
	"net"

	"github.com/bscpaz/poc-grpc-go/pb"
	"github.com/bscpaz/poc-grpc-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	listenerToPort, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	grpcServer := grpc.NewServer()

	//Attaching the implementation of AddUser to server
	//It comes from "user_grpc.pb.go" file
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listenerToPort); err != nil {
		log.Fatalf("Could not serve %v", err)
	}
}
