package main

import (
	"context"
	pb "grpc_tutorial/greet/proto"
	"log"
)


func (server *Server) Greet(ctx context.Context, request *pb.GreetRequest) (response *pb.GreetResponse, err error){
	log.Printf("Greet function was invoked  with %v\n", request)
	return &pb.GreetResponse{
		Result: "Hello " + request.FirstName,
	}, nil
}

func (server *Server) HalloThere(ctx context.Context, request *pb.GreetRequest) (response *pb.GreetResponse, err error){
	log.Printf("Greet function was invoked  with %v\n", request)
	return &pb.GreetResponse{
		Result: "Hallo There" + request.FirstName,
	},nil
}