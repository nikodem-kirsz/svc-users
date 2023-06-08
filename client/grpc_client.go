package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/nikodem-kirsz/svc-users/api/proto/users"
)

var (
	serverAddr = flag.String("addr", "localhost:3000", "The server addres in the form of host:port")
)

func main() {
	flag.Parse()
	// var opts []grpc.DialOption
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewUsersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetAllUsers(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Could not get all users from grpc server: %v", err)
	}
	log.Printf("All users: %v", r.GetUsers())

	x, err := c.GetUserById(ctx, &pb.GetUserByIdRequest{Id: "1"})
	if err != nil {
		log.Fatalf("Could not users from grpc server: %v", err)
	}

	log.Printf("User with id 1: %v", x)

}
