package main

import (
    "context"
	"strings"
	"bufio"
	"fmt"
    "log"
    "os"
    "google.golang.org/grpc"
	"client/proto"
)

func main() {
    conn, err := grpc.Dial("grpc-server:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Filed to connection: %v", err)
    }
    defer conn.Close()

    c := chat.NewChatServiceClient(conn)

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter message: ")
        text, err := reader.ReadString('\n')
        if err != nil {
            log.Fatalf("Failed to read standard input: %v", err)
        }
        text = strings.TrimSpace(text)
        if text == "exit" {
            break
        }

        response, err := c.SendMessage(context.Background(), &chat.Message{Text: text})
        if err != nil {
            log.Fatalf("Failed to send message: %v", err)
        }
        log.Printf("Server response: %s", response.Text)
    }
}
