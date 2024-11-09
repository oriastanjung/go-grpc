package main

import (
	"context"
	"fmt"
	pb "grpc_tutorial/calculator/proto"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (server *Server) SQRT(ctx context.Context, input *pb.SQRTRequest) (*pb.SQRTResponse, error){
	log.Printf("SQRT was invoked with %v", input)

	number := input.Number

	if number <0 {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("Receive the negative number %d",number),
		)
	}

	return &pb.SQRTResponse{
		Result: int32(math.Sqrt(float64(number))),
	}, nil
}