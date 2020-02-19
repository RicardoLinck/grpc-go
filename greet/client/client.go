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
