//This is the implementations of the UserServiceServer interface to attach it to server.
// See:
//   type UserServiceServer interface {
// 	  AddUser(context.Context, *User) (*User, error)
// 	  mustEmbedUnimplementedUserServiceServer()
//   }

package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/bscpaz/poc-grpc-go/pb"
)

/*
It must implement the interface below:

	type UserServiceServer interface {
		AddUser(context.Context, *User) (*User, error)
		AddUserVerbose(*User, UserService_AddUserVerboseServer) error
		AddUsers(UserService_AddUsersServer) error
	}
*/

type UserService struct {
	pb.UnimplementedUserServiceServer
}

//Added just to register this service into server.
func NewUserService() *UserService {
	return &UserService{}
}

//Implementation of "AddUser(context.Context, *User) (*User, error)"
func (*UserService) AddUser(ctx context.Context, request *pb.User) (*pb.User, error) {

	//MOCK of inserting user into database:
	fmt.Printf("Inserting '%s' into database ...", request.GetName())

	return &pb.User{
		Id:    123456,
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting into database...",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted.",
		User: &pb.User{
			Id:    12,
			Name:  req.Name,
			Email: req.Email,
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed.",
		User: &pb.User{
			Id:    12,
			Name:  req.Name,
			Email: req.Email,
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatalf("Error receiving users as stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

		fmt.Println("Adding ", req.GetName())
	}
}

func (*UserService) AddUserBiDirectionStream(stream pb.UserService_AddUserBiDirectionStreamServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error receiving stream from client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})

		if err != nil {
			log.Fatalf("Error sending stream to the client: %v", err)
		}

	}
}
