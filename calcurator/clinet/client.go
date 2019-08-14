package main

import (
	"context"
	"fmt"
	"github.com/stkhr/study-go-gRPC/calcurator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
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
	doBiDiStreaming(c)

	doError(c)
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

func doBiDiStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println("[INFO] starting bidi streaming")

	stream, err := c.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("[ERROR] failed to open stream and calling findmaximum: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 6, 125, 745, 134, 136, 96784, 1433, 146}
		for _, number := range numbers {
			fmt.Printf("sending number: %v\n", number)
			stream.Send(&calcpb.FindMaximumRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// receive
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("[ERROR] failed to read server stream: %v", err)
				break
			}
			maximum := res.GetMaximum()
			fmt.Printf("Received a new maximum of ...: %v\n", maximum)
		}
		close(waitc)
	}()
	<-waitc
}

func doError(c calcpb.CalculatorServiceClient) {
	fmt.Println("[INFO] starting example of error")
	number := int32(-25)
	res, err := c.SquareRoot(context.Background(), &calcpb.SquareRootRequest{
		Number: number,
	})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("probably sent a negative number!")
			}
		} else {
			log.Fatalf("[ERROR] failed to call SquareRoot: %v", err)
		}
	}
	fmt.Printf("Resulet of square root of %v : %v", number, res.GetNumberRoot())
}
