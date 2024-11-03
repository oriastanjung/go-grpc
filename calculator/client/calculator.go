package main

import (
	"context"
	pb "grpc_tutorial/calculator/proto"
	"io"
	"log"
)

func doPrimesManyTimes(client pb.CalculatorServiceClient, input int32){
	log.Println("do greet many time was invoked")
	request := &pb.PrimeRequest{
		Number: input,
	}
	stream, err := client.Primes(context.Background(), request)
	
	if err != nil{
		log.Fatalf("Error on streaming %v\n",err)
	}
	
	var results []int32;
	for{
		msg,err := stream.Recv()
		if err == io.EOF{
			break
		}

		if err != nil{
			log.Fatalf("Eror on reading message %v\n",err)
		}
		results= append(results, msg.Output)
	}
	
	log.Printf("Result : %v ",results)

}