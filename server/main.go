package main

import (
    "context"
    "log"
    "net"
	"google.golang.org/grpc"
	"server/proto"
)

type server struct {
    chat.UnimplementedChatServiceServer
}

func (s *server) SendMessage(ctx context.Context, message *chat.Message) (*chat.Response, error) {
    log.Printf("Received message from client: %s", message.Text)
    return &chat.Response{Text: "Message received: " + message.Text}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    s := grpc.NewServer()
    chat.RegisterChatServiceServer(s, &server{})
    log.Println("Server listening at", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
