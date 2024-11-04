package main

import (
	"fmt"
	pb "grpc_tutorial/greet/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func (server *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error{
	log.Println("Method LongGreet was inovked")
	
	res := ""

	for {
		req , err := stream.Recv()

		if err == io.EOF{
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil{
			log.Fatalf("Error on reading client stream")
		}
		log.Printf("receiving request : %v",req)
		res+= fmt.Sprintf("Hello %s!\n",req.FirstName)
	}
}
	