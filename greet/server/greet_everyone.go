package main

import (
	pb "grpc_tutorial/greet/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func (server *Server) GreetEveryone(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetResponse]) error{
	log.Println("GreetEveryone is invoke")

	for{
		request, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading request stream %v",request)
		}
		res := "Hello "+ request.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil{
			log.Fatalf("Error while send data to client %v",err)
		}
	}

}