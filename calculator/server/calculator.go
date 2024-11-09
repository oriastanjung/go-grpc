package main

import (
	"context"
	pb "grpc_tutorial/calculator/proto"
	"io"
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


func (server *Server) CalculateAvg(stream grpc.ClientStreamingServer[pb.AvgRequest, pb.AvgResponse]) error{
	log.Printf("CalculateAvg is Invoked")

	var arrayNumber []int32

	for{
		req,err := stream.Recv()
		if err == io.EOF{
			var sum int32
			for _ ,item := range arrayNumber{
				sum+=item
			}
			average := float64(sum) / float64(len(arrayNumber))
			return stream.SendAndClose(&pb.AvgResponse{
				AverageResult: average,
			})
		}
		if err != nil{
			log.Fatalf("Error on reading stream : %v\n",err)
		}
		arrayNumber = append(arrayNumber, req.Number)
	}
}

	
func (server *Server) MaxAPI(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error{
	log.Printf("MaxAPI RPC is Invoke")
	currentNumber := 0
	for{
		req,err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil{
			log.Fatalf("Error get requiest from stream %v",err)
		}

		if currentNumber < int(req.Number){
			err = stream.Send(&pb.MaxResponse{
				Result: req.Number,
			})

			if err!= nil{
				log.Fatalf("Error on sending response to stream %v", err)
			}

			currentNumber = int(req.Number)
		}
	} 
}