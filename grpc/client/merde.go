package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, ":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewChatServiceClient(conn)

	pong, err := client.SayHello(ctx, &Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pong)
}
