package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	//Building the "payload" to the request.
	req := &pb.User{
		Id:    0,
		Name:  "Bruno Paz",
		Email: "soujava@gmail.com",
	}

	//It makes the stream request.
	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	//Infitity loop to receive the steam status
	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			//There is no more stream return
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    1,
			Name:  "Jack Bauer",
			Email: "jack@mail.com",
		},
		&pb.User{
			Id:    2,
			Name:  "Nina Myers",
			Email: "nina@mail.com",
		},
		&pb.User{
			Id:    3,
			Name:  "Tony Almeida",
			Email: "toni@mail.com",
		},
		&pb.User{
			Id:    4,
			Name:  "David Palmer",
			Email: "david@mail.com",
		},
		&pb.User{
			Id:    5,
			Name:  "Chloe O'Brian",
			Email: "chloe@mail.com",
		},
		&pb.User{
			Id:    6,
			Name:  "George Mason",
			Email: "mason@mail.com",
		},
		&pb.User{
			Id:    7,
			Name:  "Michelle Dessler",
			Email: "michelle@mail.com",
		},
		&pb.User{
			Id:    8,
			Name:  "Curtis Manning",
			Email: "curtis@mail.com",
		},
		&pb.User{
			Id:    9,
			Name:  "Charles Logan",
			Email: "logan@mail.com",
		},
		&pb.User{
			Id:    10,
			Name:  "Milo Pressman",
			Email: "milo@mail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Could not sand stream request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Could not receive stream response: %v", err)
	}

	fmt.Println(res)
}
