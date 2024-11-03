package main

import (
	"context"
	pb "grpc_tutorial/calculator/proto"
	"log"

	"google.golang.org/grpc"
)


func (server *Server) Sum(ctx context.Context, request *pb.CalculatorRequest) (*pb.CalculatorResponse, error){
	log.Printf("Sum getting invoked %v", request)
	return &pb.CalculatorResponse{
		Result: request.FirstNumber + request.SecondNumber,
	},nil
}


func (server *Server)Primes(request *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error{
	log.Printf("Sum getting invoked %v", request)

	var k int32 = 2
	N := request.Number

	for N >1 {
		if N % k == 0{
			stream.Send(&pb.PrimeResponse{
				Output: k,
			})
			N = N / k
		}else {
			k+=1
		}
	}

	return nil
}