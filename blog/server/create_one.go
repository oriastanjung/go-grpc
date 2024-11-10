package main

import (
	"context"
	"fmt"
	pb "grpc_tutorial/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server)CreateOneBlog(ctx context.Context, input *pb.Blog) (*pb.BlogId, error){
	log.Printf("CreateOneBlog invoked with %v",input)

	data := BlogItem{
		AuthorId: input.AuthorId,
		Title: input.Title,
		Content: input.Content,
	}

	res,err := collection.InsertOne(ctx,data)
	if err!= nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error %v", err))
	}

	objectId , ok := res.InsertedID.(primitive.ObjectID)

	if !ok{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to ObjectID"))
	}

	return &pb.BlogId{
		Id: objectId.Hex(),
	},nil

}