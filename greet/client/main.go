package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc_tutorial/greet/proto"
)

var addr string = "localhost:27021"

func main(){

	connection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
		panic(err)
	}

	defer connection.Close()
	client := pb.NewGreetServiceClient(connection)
	// client_calculator := pb.NewCalculatorServiceClient(connection)
	doGreet(client)
	callHelloThere(client)

	// callSum(client_calculator)
	// doGreetManyTimes(client)
	// doLongGreet(client)
	doGreetEveryone(client)
}