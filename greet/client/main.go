package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "grpc_tutorial/greet/proto"
)

var addr string = "localhost:27021"

func main(){
	tls := true
	options := []grpc.DialOption{}

	if tls{
		cerFile := "ssl/ca.crt"
		
		creds,err := credentials.NewClientTLSFromFile(cerFile,"")
		if err != nil{
			log.Fatalf("Error on loading cert %v", err)
		}
		options = append(options, grpc.WithTransportCredentials(creds))
		
	}
	connection, err := grpc.Dial(addr, options...)
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
	// doGreetWithDeadlines(client,5*time.Second)
}