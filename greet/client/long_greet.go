package main

import (
	"context"
	pb "grpc_tutorial/greet/proto"
	"log"
	"time"
)

func doLongGreet(client pb.GreetServiceClient){
	log.Println("doLongGreet is invoked")


	reqs := []*pb.GreetRequest{
		{FirstName: "Orias"},
		{FirstName: "Tanjung"},
		{FirstName: "Pinang"},
	}

	stream, err := client.LongGreet(context.Background())
	if err != nil{
		log.Fatalf("Erorr while calling LongGreet %v\n",err)
	}
	for _,req := range reqs{
		log.Printf("Sending req %v ",req)
		stream.Send(&pb.GreetRequest{
			FirstName: req.FirstName,
		})

		time.Sleep(1*time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("Error while receiving response from LongGreet %v\n",err)
	}
	log.Printf("LongGreet : %s\n",res.Result)


}