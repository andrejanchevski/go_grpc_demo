package main

import (
	"io"
	"log"

	pb "github.com/andrejanchevski/go_grpc_demo/proto"
)

func (s *helloServer) SayHelloBiDirectionalStream(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Got Request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
