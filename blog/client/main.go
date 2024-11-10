package main

import (
	"log"

	pb "grpc_tutorial/blog/proto"

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

	client := pb.NewBlogServiceRoutesClient(connection)
	
	id:= createOneBlog(client)
	getOneBlog(client, "67302455534222458e75016f")
	updateData := &pb.Blog{
		AuthorId: "New Orias",
		Title: "Test 2",
		Content: "Content 1",
	}
	updateOneBlog(client,"67302455534222458e75016f",updateData)
	GetAllBlogs(client)
	deleteOneBlog(client,id)
}
