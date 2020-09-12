package main

import (
	"log"

	"github.com/kywang1/grpc_tutorial/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the clien!",
	}
	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling sayHello: %s: ", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
