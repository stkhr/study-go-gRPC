package main

import (
	"context"
	"fmt"
	"github.com/stkhr/study-go-gRPC/calcurator/calcpb"
	"google.golang.org/grpc"
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
}

func doUnary(c calcpb.CalculatorServiceClient) {
	fmt.Println("[INFO] starting unary client")
	req := &calcpb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("[ERROR] error with calling Greeting: %v", err)
	}
	log.Printf("[INFO] response: %v", res.SumResult)
}
