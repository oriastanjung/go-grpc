package main

import (
	"context"
	pb "grpc_tutorial/blog/proto"
	"log"
)

func createOneBlog(client pb.BlogServiceRoutesClient) string{
	log.Printf("CreateOneBlog invoked")

	blog := &pb.Blog{
		AuthorId: "Orias",
		Title: "My first Blog",
		Content: "Content of the first blog",
	}

	res, err := client.CreateOneBlog(context.Background(), blog)

	if err != nil{
		log.Fatalf("Error on creating blog %v",err)
	}

	log.Printf("Blog has been created : %s ", res.Id)
	return res.Id
}