package main

import (
	"context"
	pb "grpc_tutorial/calculator/proto"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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


func doMaxAPI(client pb.CalculatorServiceClient){
	log.Println("doMaxAPI invoked")
	md:=metadata.New(map[string]string{"authorization":"test-token"})
	ctx := metadata.NewOutgoingContext(context.Background(),md)
	stream,err := client.MaxAPI(ctx)
	if err != nil{
		log.Fatalf("Error while streaming %v",err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}
	var maxNumber []int32 
	waitchan := make(chan struct{})

	go func(){
		for _, req := range reqs {
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func(){
		for{
			res, err := stream.Recv()
			if err == io.EOF{
				break
			}

			if err!=nil{
				log.Fatalf("Error on receiving Stream %v", err)
				break
			}
			// log.Println(res.Result)
			maxNumber = append(maxNumber, res.Result)
		}
		close(waitchan)
	}()
	
	<-waitchan
	log.Printf("MaxNUmber : %v",maxNumber)
}


func doSqrt(client pb.CalculatorServiceClient, number int32){
	log.Println("doSqrt was invoked")

	res,err := client.SQRT(context.Background(),&pb.SQRTRequest{
		Number: number,
	})

	if err!=nil{
		erorr, ok := status.FromError(err)

		if ok{
			log.Printf("message from server %s",erorr.Message())
			log.Printf("error code from server %s",erorr.Code())

			if erorr.Code() == codes.InvalidArgument{
				log.Println("We probably send negative number")
				return
			}
		}else{
			log.Fatalf("A non GRPC Error : %v",err)
		}
	}

	log.Printf("sqrt : %v",res.Result)
}