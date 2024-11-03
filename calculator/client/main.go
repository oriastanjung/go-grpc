package main

import (
	"context"
	"log"

	pb "grpc_tutorial/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:2727"
func main(){
	connection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
		panic(err)
	}
	defer connection.Close()

	client := pb.NewCalculatorServiceClient(connection)
	doSum(client)
	doPrimesManyTimes(client,120)
}

func doSum(client pb.CalculatorServiceClient){

	log.Println("doSum was invoked")

	response,err := client.Sum(context.Background(), &pb.CalculatorRequest{
		FirstNumber: 10,
		SecondNumber: 3,
	})

	if err != nil{
		log.Fatalf("Couldnt Sum %v\n", err)
	}

	log.Printf("Sum: %d\n", response.Result)
}