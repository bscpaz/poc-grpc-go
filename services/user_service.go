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

	"github.com/bscpaz/poc-grpc-go/pb"
)

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
	fmt.Println("Inserting '%s' into database ...", request.GetName())

	return &pb.User{
		Id:    123456,
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}
