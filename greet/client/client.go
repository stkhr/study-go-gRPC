package main

import (
	"context"
	"fmt"
	"github.com/stkhr/study-go-gRPC/greet/greetpb"
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

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

	doServerStreaming(c)
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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("[INFO] starting server streaming client")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "test2",
			LastName:  "hoge2",
		},
	}
	res, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("[ERROR] error with calling GreetManyTimes: %v", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			// reached the end of stream
			break
		} else if err != nil {
			log.Fatalf("[ERROR] error with reading stream: %v", err)
		} else {
			log.Printf("[INFO] response: %v", msg.GetResult())
		}
	}

}
