package main

import (
	"context"
	"fmt"
	"log"

	"github.com/andres07a/grpc-practice/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client: Hi!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {

	log.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{
			NumOne: 10,
			NumTwo: 3,
		},
	}
	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calculator RPC: %v", err)
	}
	log.Printf("Response from Greet: %v\n", res.Result)
}
