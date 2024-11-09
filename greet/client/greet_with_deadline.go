package main

import (
	"context"
	pb "grpc_tutorial/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func doGreetWithDeadlines(client pb.GreetServiceClient, timeout time.Duration){
	log.Println("DoGreetWithDeadline is Invoked")

	ctx, cancel := context.WithTimeout(context.Background(),timeout)
	defer cancel()

	req:= &pb.GreetRequest{
		FirstName: "Orias",
	}
	res, err := client.GreetWithDeadline(ctx,req)

	if err != nil{
		e,ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded{
				log.Println("Deadline Exceeded")
				return
			}
			log.Fatalf("Unexpected GRPC Error %v",e)
		}else{
			log.Printf("Error %v",err)
		}
	}

	log.Printf("GreetWithDeadline Response : %v", res.Result)
}