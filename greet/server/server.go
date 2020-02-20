package main

import (
	"context"
	"fmt"
	"greet/greetpb"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("Greet rpc invoked!")
	first := req.Greeting.FirstName
	return &greetpb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", first),
	}, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Println("GreetManyTimes rpc invoked!")
	first := req.Greeting.FirstName

	wg := &sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(stream greetpb.GreetService_GreetManyTimesServer, i int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			stream.Send(&greetpb.GreetManyTimesResponse{
				Result: fmt.Sprintf("Hello %s number %d", first, i),
			})
			wg.Done()
		}(stream, i)
	}

	wg.Wait()

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Println("LongGreet rpc invoked!")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		firstName := req.Greeting.FirstName
		result += fmt.Sprintf("Helo %s! ", firstName)
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
