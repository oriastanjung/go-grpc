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



func doAverage(client pb.CalculatorServiceClient){
	log.Printf("doAverage is invoked")

	requests := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}
	stream, err := client.CalculateAvg(context.Background())
	if err != nil {
		log.Fatalf("Error on reading stream %v\n",err)
	}
	for _, req := range requests{
		stream.Send(&pb.AvgRequest{
			Number: req.Number,
		})
	}

	res,err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("Error on getting response %v\n",err)
	}

	log.Printf("Average from %v is %v",requests,res.AverageResult)
}