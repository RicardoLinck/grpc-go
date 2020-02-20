package main

import (
	"calculator/calculatorpb"
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Sum rpc invoked with req: %v\n", req)
	return &calculatorpb.SumResponse{
		Result: req.NumA + req.NumB,
	}, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Printf("PrimeNumberDecomposition rpc invoked with req: %v\n", req)
	k := int32(2)
	n := req.Input

	for n > 1 {
		if n%k == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				Result: k,
			})
			n = n / k
		} else {
			k++
		}
	}

	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	log.Printf("PrimeNumberDecomposition rpc invoked with stream: %v\n", stream)
	count := 0
	sum := float64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: sum / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		sum += req.Input
		count++
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	log.Printf("FindMaximum rpc invoked with stream: %v\n", stream)
	max := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		if req.Input > max {
			max = req.Input
			stream.Send(&calculatorpb.FindMaximumResponse{Result: max})
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
