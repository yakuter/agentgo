package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"

	"github.com/yakuter/agentgo/pb"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Executer interface {
	Execute() (string, error)
}

type Linux struct {
	App  string
	Args []string
}

func (l *Linux) Execute() (string, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(l.App, l.Args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		return outStr, fmt.Errorf(errStr)
	}

	return outStr, nil
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedCommandServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Send(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error) {

	// Log incoming command
	log.Printf("\nApplication: %v\nArguments: %v", in.GetApp(), in.GetArgs())

	// Define Executer interface
	var e Executer

	// Assign a Linux command to Executer
	e = &Linux{in.GetApp(), in.GetArgs()}

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

func main() {

	// Define listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create grpc server
	s := grpc.NewServer()

	// Register server struct to grpc server
	pb.RegisterCommandServiceServer(s, &server{})

	// Start grpc server
	log.Printf("Agent started listening localhost%s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
