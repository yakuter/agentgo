package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/yakuter/agentgo/agentgo-server/internal/config"
	"github.com/yakuter/agentgo/pb"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	// Create a FlagSet and sets the usage
	fs := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	// Configure the options from the flags/config file
	opts, err := config.ConfigureOptions(fs, os.Args[1:])
	if err != nil {
		config.PrintUsageErrorAndDie(err)
	}

	// If -help flag is defined, print help
	if opts.ShowHelp {
		config.PrintHelpAndDie()
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create and register a new client
	c := pb.NewCommandServiceClient(conn)

	// Create the context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Generate app and arguments to send
	command := pb.CommandRequest{
		App:  opts.App,
		Args: opts.Args,
	}

	// Send command
	r, err := c.Send(ctx, &command)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Command Result:\n%s", r.GetResult())
}
