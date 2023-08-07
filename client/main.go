package main

import (
	"log"

	pb "github.com/andrejanchevski/go_grpc_demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection issue: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := pb.NamesList{
	// 	Names: []string {"Andrej", "Alice", "Bob"}
	// }

	callSayHello(client)
}
