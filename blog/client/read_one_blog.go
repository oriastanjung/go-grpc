package main

import (
	"context"
	pb "grpc_tutorial/blog/proto"
	"log"
)


func getOneBlog(client pb.BlogServiceRoutesClient, id string) *pb.Blog{
	log.Printf("getOneBlog invoked with id : %v", id)
	input := &pb.BlogId{
		Id: id,
	}
	res,err := client.GetOneBlog(context.Background(), input)

	if err != nil{
		log.Fatalf("Error : %v",err)
	}

	log.Printf("Blog get is : %v",res)

	return res
}