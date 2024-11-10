package main

import (
	"context"
	pb "grpc_tutorial/blog/proto"
	"log"
)

func deleteOneBlog(client pb.BlogServiceRoutesClient, id string){
	log.Printf("delete one blog for %v ", id)
	input := &pb.BlogId{
		Id: id,
	}
	_, err := client.DeleteOneBlog(context.Background(),input)
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("Delete Success")

}