package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bscpaz/poc-grpc-go/pb"
	"google.golang.org/grpc"
)

func main() {
	//It creates a gRPC connection to a server
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connection to gRPC server: %v", err)
	}

	//This will close the connection when "connection" variable is not used anymore.
	defer connection.Close()

	//You get all implemented methods to server "for free" (it cames from the stubs).
	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	//Building the "payload" to the request.
	req := &pb.User{
		Id:    0,
		Name:  "Bruno Paz",
		Email: "soujava@gmail.com",
	}

	//It makes the request.
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}
