package main

import (
	"calculator/calculatorpb"
	"context"
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

	c := calculatorpb.NewCalculatorServiceClient(cc)
	sum(c)
	primeNumberDecomposition(c)
}

func sum(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		NumA: 3,
		NumB: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Sum RPC: %v", err)
	}
	log.Printf("Response: %d\n", res.Result)
}

func primeNumberDecomposition(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Input: 54224620,
	}
	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading PrimeNumberDecomposition stream: %v", err)
		}
		log.Println(res.Result)
	}
}
