package main

import (
	"context"
	pb "grpc_tutorial/greet/proto"
	"log"
)

func callSum(client pb.CalculatorServiceClient){
	log.Println("callSum was invoked")

	response,err := client.Sum(context.Background(), &pb.CalculatorRequest{
		FirstNumber: 10,
		SecondNumber: 3,
	})

	if err != nil{
		log.Fatalf("Couldnt Sum %v\n", err)
	}

	log.Printf("Sum: %d\n", response.Result)
}