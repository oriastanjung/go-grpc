package main

import (
	"fmt"
	pb "grpc_tutorial/greet/proto"
	"log"

	"google.golang.org/grpc"
)

func (server *Server) GreetManyTimes(req *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("Greet many time was invoked : %v\n",req)
	for i := 0 ; i <10; i++ {
		res := fmt.Sprintf("Hello %s number %d ", req.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}