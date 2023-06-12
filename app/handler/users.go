package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/nikodem-kirsz/svc-users/api/proto/users"
	"github.com/nikodem-kirsz/svc-users/app/db/model"
)

type User struct {
	UUID         string
	email        string
	first_name   string
	last_name    string
	phone_number string
	locale       string
}

type UsersHandler struct {
	pb.UnimplementedUsersServer
	storage model.Storage
	users   []*User
}

func New(storage model.Storage) (*UsersHandler, error) {
	s := []*User{
		{UUID: "1", email: "niko@niko.pl", first_name: "niko", last_name: "okin", phone_number: "1234567", locale: "en-GB"},
		{UUID: "2", email: "niko2@niko.pl", first_name: "niko2", last_name: "okin2", phone_number: "1234567", locale: "en-GB"},
		{UUID: "3", email: "niko3@niko.pl", first_name: "niko3", last_name: "okin3", phone_number: "1234567", locale: "en-GB"},
	}
	return &UsersHandler{
		storage: storage,
		users:   s,
	}, nil
}

func (s *UsersHandler) GetAllUsers(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllUsersResponse, error) {
	users := []*pb.User{}

	for _, user := range s.users {
		users = append(users, s.buildUserResponse(ctx, user))
	}

	usersResponse := pb.GetAllUsersResponse{
		Users: users,
	}

	return &usersResponse, nil
}

func (s *UsersHandler) GetUser(context.Context, *emptypb.Empty) (*pb.UserResponse, error) {

	return &pb.UserResponse{}, nil
}

func (s *UsersHandler) GetUserById(ctx context.Context, filter *pb.GetUserByIdRequest) (*pb.UserResponse, error) {
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

func (s *UsersHandler) RegisterUser(context.Context, *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return &pb.RegisterUserResponse{}, nil
}

func (s *UsersHandler) UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *UsersHandler) buildUserResponse(ctx context.Context, user *User) *pb.User {
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
