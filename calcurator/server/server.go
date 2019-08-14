package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"

	"github.com/stkhr/study-go-gRPC/calcurator/calcpb"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("[DEBUG] request: %v \n", req)
	firstNum := req.FirstNumber
	secondNum := req.SecondNumber
	sum := firstNum + secondNum
	res := &calcpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calcpb.PrimeNumberDecompositionRequest, stream calcpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("[DEBUG] stream request: %v \n", req)

	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calcpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Println("[INFO] divisor has increased to %v", divisor)
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calcpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("[DEBUG] compute average request\n")
	sum := int32(0)
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&calcpb.ComputeAveragResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("[ERROR] reading client stream: %v", err)
		}
		sum += req.GetNumber()
		count++
	}
}

func (*server) FindMaximum(stream calcpb.CalculatorService_FindMaximumServer) error {
	fmt.Printf("[DEBUG] find maximum request\n")
	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("[ERROR] failed to read client stream: %v", err)
		}
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			err = stream.Send(&calcpb.FindMaximumResponse{
				Maximum: maximum,
			})
			if err != nil {
				log.Fatalf("[ERROR] failed to send data to clinet: %v", err)
			}
		}
	}
}

func main() {
	fmt.Println("[INFO] start calc server")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
