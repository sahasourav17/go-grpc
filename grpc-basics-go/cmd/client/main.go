package main

import (
	"context"
	"log"
	"time"

	greeterv1 "example.com/grpc-basics-go/gen/greeter/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}
	defer conn.Close()
	client := greeterv1.NewGreeterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	resp, err := client.SayHello(ctx, &greeterv1.SayHelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("SayHello response: %v", resp.GetMessage())
}
