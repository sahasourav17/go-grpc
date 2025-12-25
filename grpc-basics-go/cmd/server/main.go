package main

import (
	"context"
	"fmt"
	"log"
	"net"

	greeterv1 "example.com/grpc-basics-go/gen/greeter/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type greeterServer struct {
	greeterv1.UnimplementedGreeterServiceServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *greeterv1.SayHelloRequest) (*greeterv1.SayHelloResponse, error) {
	name := req.GetName()
	if name == "" {
		name = "World"
	}
	return &greeterv1.SayHelloResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	greeterv1.RegisterGreeterServiceServer(grpcServer, &greeterServer{})
	reflection.Register(grpcServer)

	log.Println("gRPC server listening on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
