package main

import (
	"context"
	"fmt"
	"github.com/stkhr/study-go-gRPC/greet/greetpb"
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

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("[INFO] starting unary client")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "test",
			LastName:  "hoge",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("[ERROR] error with calling Greeting: %v", err)
	}
	log.Printf("[INFO] response: %v", res.Result)
}
