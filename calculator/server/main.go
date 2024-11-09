package main

import (
	"log"
	"net"

	pb "grpc_tutorial/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:2727"

type Server struct {
    pb.CalculatorServiceServer
}

func main(){
    listener,err := net.Listen("tcp",addr)
    if err != nil{
        log.Printf("Error on Listening %v\n",err)
    }

    defer listener.Close()
	log.Printf("listening on %s\n", addr)

    options:= []grpc.ServerOption{}
    options = append(options, grpc.UnaryInterceptor(UnaryMetadataInterceptor))
    options = append(options, grpc.StreamInterceptor(StreamMetadataInterceptor))

    serverInstance := grpc.NewServer(options...)

    pb.RegisterCalculatorServiceServer(serverInstance, &Server{})
    reflection.Register(serverInstance)
    if err = serverInstance.Serve(listener); err != nil{
        log.Fatalf("Failed on Serve %v\n",err)
    }
}