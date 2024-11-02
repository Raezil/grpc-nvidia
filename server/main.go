package main

import (
	"log"
	"net"
	pb "nvidia"

	"google.golang.org/grpc"
)

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the NVIDIA server
	pb.RegisterNVIDIAServer(grpcServer, &pb.NvidiaServer{})

	log.Println("NVIDIA gRPC server is running on port 50051...")

	// Start serving
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
