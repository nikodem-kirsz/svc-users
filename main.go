package main

import (
	"flag"
	"fmt"
	"net"

	pb "github.com/nikodem-kirsz/svc-users/api/proto/users"
	"github.com/nikodem-kirsz/svc-users/app"
	"github.com/nikodem-kirsz/svc-users/app/configuration"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 3000, "Port where server listens too")
)

func main() {
	flag.Parse()
	logger := log.NewEntry(log.StandardLogger())
	logger.Println("Loading configuration...")

	conf, err := configuration.Load()
	if err != nil {
		logger.Fatalf("Fail to load configuration for the environment")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		logger.Fatalf("Fail to listen %v", err)
	}

	if err != nil {
		logger.Fatal("unable to configure service, %w", err)
	}

	logger.Printf("Initializing application...")

	app, err := app.BuildDependencies(conf, logger)
	if err != nil {
		logger.Fatalf("failed to build application dependencies")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUsersServer(grpcServer, app.Handler)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
