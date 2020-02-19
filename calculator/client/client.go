package main

import (
	"calculator/calculatorpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewSumServiceClient(cc)
	req := &calculatorpb.SumRequest{
		NumA: 3,
		NumB: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling RPC: %v", err)
	}
	fmt.Printf("Response: %d\n", res.Result)
}
