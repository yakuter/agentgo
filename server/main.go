package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/yakuter/agentgo/pb"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var app string
	var args arrayFlags

	flag.StringVar(&app, "app", "", "Application to execute command: ls")
	flag.Var(&args, "arg", "Arguments for application: -lah")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create and register a new client
	c := pb.NewCommandServiceClient(conn)

	// Create the context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Generate app and arguments to send
	command := pb.CommandRequest{
		App:  app,
		Args: args,
	}

	// Send command
	r, err := c.Send(ctx, &command)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Command Result:\n%s", r.GetResult())
}
