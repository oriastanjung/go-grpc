package main

import (
	"context"
	pb "grpc_tutorial/blog/proto"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func GetAllBlogs(client pb.BlogServiceRoutesClient){
	log.Printf("Call GetAllBlogs")
	stream,err := client.GetAllBlogs(context.Background(),&emptypb.Empty{})
	if err!=nil{
		log.Fatalf("Error while reading stream")
	}
	blogs:=[]*pb.Blog{}
	for{
		res,err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatal(err)
			break
		}
		blogs = append(blogs, res)
			
	}

	log.Printf("Blogs List is : \n")
	log.Printf("No. - ID - Author - Title - Content")
	for idx, item := range blogs{
		log.Printf("%d. - %v - %v - %v - %v",idx,item.Id, item.AuthorId, item.Title, item.Content)
	}
}