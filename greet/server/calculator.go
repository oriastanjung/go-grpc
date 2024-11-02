package main

import (
	"context"
	pb "grpc_tutorial/greet/proto"
	"log"
)


func (server *CalculatorServer) Sum(ctx context.Context, request *pb.CalculatorRequest) (response *pb.CalculatorResponse, err error){
	log.Printf("Sum function was invoked  with %v\n", request)
	return &pb.CalculatorResponse{
		Result: request.FirstNumber + request.SecondNumber,
	},nil
}