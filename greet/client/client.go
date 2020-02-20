package main

import (
	"context"
	"greet/greetpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	greet(c)
	greetManyTimes(c)
	longGreet(c)
}

func greet(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ricardo",
			LastName:  "Linck",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Greet RPC: %v", err)
	}
	log.Printf("Response: %s\n", res.Result)
}

func greetManyTimes(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ricardo",
			LastName:  "Linck",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GreetManyTimes RPC: %v", err)
	}

	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading GreetManyTimes stream: %v", err)
		}
		log.Printf("Response: %s\n", res.Result)
	}
}

func longGreet(c greetpb.GreetServiceClient) {
	reqs := []*greetpb.LongGreetRequest{
		{Greeting: &greetpb.Greeting{
			FirstName: "Ricardo",
			LastName:  "Linck",
		}},
		{Greeting: &greetpb.Greeting{
			FirstName: "Carolina",
			LastName:  "Pacheco",
		}},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error calling LongGreet RPC: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet RPC: %v", err)
	}

	log.Printf("Response: %s\n", response)

}
