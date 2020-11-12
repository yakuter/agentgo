package api

import (
	"context"
	"log"

	"github.com/yakuter/agentgo/agentgo-agent/internal/app"
	"github.com/yakuter/agentgo/pb"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedCommandServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Send(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error) {

	// Log incoming command
	log.Printf("\nApplication: %v\nArguments: %v", in.GetApp(), in.GetArgs())

	// Define Executer interface
	var e app.Executer

	// Assign a Linux command to Executer
	e = &app.Linux{in.GetApp(), in.GetArgs()}

	// Execute command and retrieve result
	result, err := e.Execute()
	if err != nil {
		log.Println(err)
	}

	// Generate grpc response
	response := &pb.CommandResponse{
		Result: result,
	}

	return response, err
}
