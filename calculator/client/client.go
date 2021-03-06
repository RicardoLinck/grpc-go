package main

import (
	"calculator/calculatorpb"
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	sum(c)
	// primeNumberDecomposition(c)
	// computeAverage(c)
	findMaximum(c)
}

func sum(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		NumA: 3,
		NumB: 10,
	}
	ctx := metadata.AppendToOutgoingContext(context.Background(), "user", "test")
	res, err := c.Sum(ctx, req)
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

func computeAverage(c calculatorpb.CalculatorServiceClient) {
	reqs := []*calculatorpb.ComputeAverageRequest{
		{Input: 1},
		{Input: 2},
		{Input: 3},
		{Input: 4},
	}
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error calling ComputeAverage RPC: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from ComputeAverage RPC: %v", err)
	}

	log.Printf("Response: %s\n", response)
}

func findMaximum(c calculatorpb.CalculatorServiceClient) {
	reqs := []*calculatorpb.FindMaximumRequest{
		{Input: 1},
		{Input: 5},
		{Input: 3},
		{Input: 6},
		{Input: 2},
		{Input: 20},
	}
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error calling ComputeAverage RPC: %v", err)
	}

	wc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			response, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving response from ComputeAverage RPC: %v", err)
				break
			}

			log.Printf("Response: %s\n", response)
		}
		close(wc)
	}()

	<-wc

}
