package main

import (
	"context"
	pb "grpc_tutorial/calculator/proto"
	"log"
)


func (server *Server) Sum(ctx context.Context, request *pb.CalculatorRequest) (*pb.CalculatorResponse, error){
	log.Printf("Sum getting invoked %v", request)
	return &pb.CalculatorResponse{
		Result: request.FirstNumber + request.SecondNumber,
	},nil
}