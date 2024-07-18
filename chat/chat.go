package chat

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}


func (s *Server) Send(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received: %v", in.GetText())
	return &Message{Text: "Hello " + in.GetUser()}, nil
}



