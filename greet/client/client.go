package main

import (
	"context"
	"fmt"
	"github.com/stkhr/study-go-gRPC/greet/greetpb"
	"google.golang.org/grpc"
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

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

	doServerStreaming(c)

	doClientStreaming(c)
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
	fmt.Println("[INFO] starting server streaming")

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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("[INFO] starting client streaming")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "fuga1",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "fuga2",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "fuga3",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "fuga4",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "fuga5",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "fuga6",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("[ERROR] error with calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Println("[INFO] Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("[ERROR] error with recieving response: %v", err)
	}
	fmt.Printf("[INFO] LongGreet Response: %v\n", res)

}
