package main

import (
	"log"
	"net"

	pb "grpc_tutorial/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	options:= []grpc.ServerOption{}
	tls:=true // change to false if needed

	if tls{
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds,err := credentials.NewServerTLSFromFile(certFile,keyFile)
		if err != nil{
			log.Fatalf("Failed login certificate %v", err)
		}
		options= append(options,grpc.Creds(creds))
	}

	serverInstance := grpc.NewServer(options...)

	pb.RegisterGreetServiceServer(serverInstance, &Server{})
	pb.RegisterCalculatorServiceServer(serverInstance, &CalculatorServer{})
	
	if err = serverInstance.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
		panic(err)
	}




}