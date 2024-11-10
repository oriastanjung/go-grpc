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


func (server *Server) UpdateOneBlog(ctx context.Context, input *pb.Blog) (*emptypb.Empty, error){
	log.Printf("UpdateOneBlog invoked with : %v ", input)

	id := input.Id
	oid,err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil,status.Errorf(codes.Internal, fmt.Sprintf("Error on parse id"))
	}

	data := &BlogItem{
		AuthorId: input.AuthorId,
		Title: input.Title,
		Content: input.Content,
	}

	res,err := collection.UpdateOne(ctx, bson.M{"_id" : oid}, bson.M{"$set":data})
	if err != nil{
		return nil, status.Errorf(codes.Internal, "Erorr on update one")
	}
	if res.MatchedCount == 0{
		return nil, status.Errorf(codes.NotFound, "Id not found")
	}

	return &emptypb.Empty{},nil
}