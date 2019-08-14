package main

import (
	"context"
	"fmt"
	"github.com/stkhr/study-go-gRPC/calcurator/calcpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("gRPC client start")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalculatorServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	fmt.Println("[INFO] starting unary client")
	req := &calcpb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("[ERROR] error with calling sum: %v", err)
	}
	log.Printf("[INFO] response: %v", res.SumResult)
}

func doServerStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println("[INFO] starting server streaming client")
	req := &calcpb.PrimeNumberDecompositionRequest{
		Number: 1234567890,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("[ERROR] error with prime number decomposition : %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("[ERROR] failed to connect server streaming rpc")
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println("[INFO] starting client streaming")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("[ERROR] failed to connect client streaming rpc: %v", err)
	}

	numbers := []int32{3, 5, 7, 11}

	for _, number := range numbers {
		fmt.Printf("sending number: %d\n", number)
		stream.Send(&calcpb.ComputeAverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("[ERROR] recieving response: %v", err)
	}

	fmt.Printf("The Average is: %v", res.GetAverage())
}
