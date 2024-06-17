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
    fmt.Println("クライアントを開始します...")
    conn, err := grpc.Dial("grpc-server:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("接続に失敗しました: %v", err)
    }
    defer conn.Close()

    c := chat.NewChatServiceClient(conn)

    reader := bufio.NewReader(os.Stdin)
    fmt.Println("入力待ちの状態です...")
    for {
        fmt.Print("Enter message: ")
        text, err := reader.ReadString('\n')
        if err != nil {
            log.Fatalf("標準入力の読み取りに失敗しました: %v", err)
        }
        text = strings.TrimSpace(text)
        if text == "exit" {
            break
        }

        response, err := c.SendMessage(context.Background(), &chat.Message{Text: text})
        if err != nil {
            log.Fatalf("メッセージ送信に失敗しました: %v", err)
        }
        log.Printf("サーバーからの応答: %s", response.Text)
    }
}

