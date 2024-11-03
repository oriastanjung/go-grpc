package main

import (
	"context"
	pb "grpc_tutorial/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(client pb.GreetServiceClient){
	log.Println("do greet many time was invoked")

	request := &pb.GreetRequest{
		FirstName: "Oriastanjung",
	}

	stream,err := client.GreetManyTimes(context.Background(),request)

	if err != nil{
		log.Fatalf("error while calling GreetManyTimes")
	}

	for{
		msg, err := stream.Recv()

		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("Error while reading stream %v\n", err)
		}

		log.Printf("GreetManyTimes : %v\n",msg.Result)
	}
}