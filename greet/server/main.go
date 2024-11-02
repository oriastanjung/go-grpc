package main

import (
	"log"
	"net"

	pb "grpc_tutorial/greet/proto"

	"google.golang.org/grpc"
)


var addr string = "0.0.0.0:27021"

type Server struct{
	pb.GreetServiceServer
}

type CalculatorServer struct {
	pb.CalculatorServiceServer
}
func main(){
	listener,err := net.Listen("tcp",addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		panic(err)
	}

	log.Printf("listening on %s\n", addr)

	serverInstance := grpc.NewServer()

	pb.RegisterGreetServiceServer(serverInstance, &Server{})
	pb.RegisterCalculatorServiceServer(serverInstance, &CalculatorServer{})
	
	if err = serverInstance.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
		panic(err)
	}




}