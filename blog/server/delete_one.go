package main

import (
	"context"
	"fmt"
	pb "grpc_tutorial/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)


func (server *Server) DeleteOneBlog(ctx context.Context, input *pb.BlogId) (*emptypb.Empty, error){
	log.Printf("Delete one blog invoked with id : %v", input.Id)

	oid, err := primitive.ObjectIDFromHex(input.Id)
	if err!= nil{
		return nil, status.Errorf(codes.Internal, "Error while parsing objectId")
	}

	filter:= bson.M{"_id":oid}
	res,err := collection.DeleteOne(ctx, filter)

	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error while deleting %v",err))

	}

	if res.DeletedCount == 0{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error Id blog not found"))
	}

	return &emptypb.Empty{},nil
	
}