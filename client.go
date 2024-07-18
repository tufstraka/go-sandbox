package main

import (
	"log"

	"github.com/tufstraka/go-sandbox/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.Send(context.Background(), &chat.Message{User: "tufstraka", Text: "Hello"})
	if err != nil {
		log.Fatalf("Failed to send: %v", err)
	}

	log.Printf("Response: %v", response.GetText())
}