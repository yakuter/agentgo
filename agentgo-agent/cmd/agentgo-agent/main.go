package main

import (
	"log"
	"net"

	"github.com/yakuter/agentgo/agentgo-agent/internal/api"
	"github.com/yakuter/agentgo/pb"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	// Define listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create grpc server
	s := grpc.NewServer()

	// Register server struct to grpc server
	api := &api.Server{}
	pb.RegisterCommandServiceServer(s, api)

	// Start grpc server
	log.Printf("Agent started listening localhost%s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
