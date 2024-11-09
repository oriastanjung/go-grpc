package main

import (
	"context"
	pb "grpc_tutorial/greet/proto"
	"io"
	"log"
	"time"
)


func doGreetEveryone(client pb.GreetServiceClient){
	log.Println("doGreetEveryone is invoked")

	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream %v")
	}

	reqs := []*pb.GreetRequest{
		{
			FirstName: "orias",
		},
		{
			FirstName: "omi",
		},
		{
			FirstName: "ate",
		},
	}

	waitc := make(chan struct{})

	go func(){
		for _, req := range reqs{
			log.Printf("Sending request %v",req)
			stream.Send(req)
			time.Sleep(1*time.Second)
		}
		stream.CloseSend()
	}()

	go func(){
		for{
			res, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil{
				log.Printf("Error on receiveing from stream %v",err)
				break
			}

			log.Printf("Received : %v\n",res.Result)
		}
		close(waitc)
	}()


	<-waitc
}