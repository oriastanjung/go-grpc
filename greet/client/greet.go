package main

import (
	"context"
	pb "grpc_tutorial/greet/proto"
	"log"
)


func doGreet(client pb.GreetServiceClient){
	log.Println("doGreet was invoked")

	response,err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "orias",
	})

	if err != nil{
		log.Fatalf("Couldnt Greet %v\n", err)
	}

	log.Printf("Greeting: %s\n", response.Result)

}


func callHelloThere(client pb.GreetServiceClient){
	log.Println("callHelloThere was invoked")

	response,err := client.HalloThere(context.Background(), &pb.GreetRequest{
		FirstName: "orias",
	})

	if err != nil{
		log.Fatalf("Couldnt Greet %v\n", err)
	}

	log.Printf("Greeting: %s\n", response.Result)

}