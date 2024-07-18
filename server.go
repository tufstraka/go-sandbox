package main

import(
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/tufstraka/go-sandbox/chat"
)


func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}