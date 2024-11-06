package main

import (
	"context"
	"fmt"
	"net"

	"myproject/userpb"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	users := map[int32]string{
		1: "Alice",
		2: "Bob",
	}

	name, exists := users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return &userpb.UserResponse{
		Id:   req.Id,
		Name: name,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("Failed to listen on port 50051: " + err.Error())
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	fmt.Println("UserService is running on port 50051")
	if err := s.Serve(listener); err != nil {
		panic("Failed to serve: " + err.Error())
	}
}
