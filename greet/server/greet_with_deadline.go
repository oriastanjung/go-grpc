package main

import (
	"context"
	pb "grpc_tutorial/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (server *Server) GreetWithDeadline(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error){
	log.Printf("GreetWithDeadline invoked %v", request)
	for i:=0;i<3;i++{
		if ctx.Err() == context.DeadlineExceeded{
			log.Println("The client canceled the request")
			return nil, status.Error(codes.Canceled,"The Client Canceled the request")
		}
		time.Sleep(1*time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello "+ request.FirstName,
	},nil
}