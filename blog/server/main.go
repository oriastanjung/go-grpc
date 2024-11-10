package main

import (
	"context"
	"log"
	"net"

	pb "grpc_tutorial/blog/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
var collection *mongo.Collection 
var addr string = "0.0.0.0:2727"

type Server struct {
    pb.BlogServiceRoutesServer
}

func main(){
	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err !=nil{
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil{
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

    listener,err := net.Listen("tcp",addr)
    if err != nil{
        log.Printf("Error on Listening %v\n",err)
    }

    defer listener.Close()
	log.Printf("listening on %s\n", addr)

    

    serverInstance := grpc.NewServer()

    pb.RegisterBlogServiceRoutesServer(serverInstance, &Server{})
    reflection.Register(serverInstance)
    if err = serverInstance.Serve(listener); err != nil{
        log.Fatalf("Failed on Serve %v\n",err)
    }
}