package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/nikodem-kirsz/svc-users/api/proto/users"
)

var (
	port = flag.Int("port", 3000, "The server port")
)

type User struct {
	UUID         string
	email        string
	first_name   string
	last_name    string
	phone_number string
	locale       string
}

type usersServer struct {
	pb.UnimplementedUsersServer
	mu    sync.Mutex
	users []*User
}

func (s *usersServer) buildUserResponse(ctx context.Context, user *User) *pb.User {
	res := &pb.User{
		Id:          user.UUID,
		Email:       user.email,
		FirstName:   user.first_name,
		LastName:    user.last_name,
		PhoneNumber: user.phone_number,
		Locale:      user.locale,
	}
	return res
}

func (s *usersServer) GetAllUsers(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllUsersResponse, error) {
	users := []*pb.User{}

	for _, user := range s.users {
		users = append(users, s.buildUserResponse(ctx, user))
	}

	usersResponse := pb.GetAllUsersResponse{
		Users: users,
	}

	return &usersResponse, nil
}

func (s *usersServer) GetUser(context.Context, *emptypb.Empty) (*pb.UserResponse, error) {

	return &pb.UserResponse{}, nil
}

func (s *usersServer) GetUserById(ctx context.Context, filter *pb.GetUserByIdRequest) (*pb.UserResponse, error) {
	for _, user := range s.users {
		if user.UUID == filter.Id {
			return &pb.UserResponse{
				Id:          user.UUID,
				Email:       user.email,
				FirstName:   user.first_name,
				LastName:    user.last_name,
				Locale:      user.locale,
				PhoneNumber: user.phone_number,
			}, nil
		}
	}
	return &pb.UserResponse{}, nil
}

func (s *usersServer) RegisterUser(context.Context, *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return &pb.RegisterUserResponse{}, nil
}

func (s *usersServer) UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	s := &usersServer{
		users: []*User{
			{UUID: "1", email: "niko@niko.pl", first_name: "niko", last_name: "okin", phone_number: "1234567", locale: "en-GB"},
			{UUID: "2", email: "niko2@niko.pl", first_name: "niko2", last_name: "okin2", phone_number: "1234567", locale: "en-GB"},
			{UUID: "3", email: "niko3@niko.pl", first_name: "niko3", last_name: "okin3", phone_number: "1234567", locale: "en-GB"},
		},
	}
	pb.RegisterUsersServer(grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
