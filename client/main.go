// client.go
package main

import (
	"context"
	"log"

	pb "nvidia" // Adjust the import path based on your project structure

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNVIDIAClient(conn)

	// Prepare the request
	req := &pb.NVIDIARequest{
		Model:  "meta/llama-3.2-3b-instruct",
		Prompt: "Tell me more about golang",
	}

	// Set a timeout for the request
	ctx := context.Background()

	// Call the Generate method
	resp, err := client.Generate(ctx, req)
	if err != nil {
		log.Fatalf("Could not generate: %v", err)
	}

	log.Printf("Response from server: %s", resp.Result)
}
