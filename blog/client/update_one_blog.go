package main

import (
	"context"
	pb "grpc_tutorial/blog/proto"
	"log"
)


func updateOneBlog(client pb.BlogServiceRoutesClient, id string, updated *pb.Blog){
	log.Printf("Update One Blog invoked with : %v",id)
	blog := &pb.Blog{
		Id: id,
		AuthorId: updated.AuthorId,
		Title: updated.Title,
		Content: updated.Content,
	}
	_,err := client.UpdateOneBlog(context.Background(), blog)
	if err != nil{
		log.Fatal(err)
	}

	log.Printf("Updated!")
}