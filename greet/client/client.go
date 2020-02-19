package main

import (
	"context"
	"fmt"
	"greet/greetpb"
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
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ricardo",
			LastName:  "Linck",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling RPC: %v", err)
	}
	fmt.Printf("Response: %s\n", res.Result)
}
