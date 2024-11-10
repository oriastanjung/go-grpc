package main

import (
	"context"
	pb "grpc_tutorial/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetOneBlog(ctx context.Context, input *pb.BlogId) (*pb.Blog, error){
	log.Printf("GetOneBlog invoked with : %v", input)
	objectId, err := primitive.ObjectIDFromHex(input.Id)
	if err != nil{
		return nil, status.Errorf(
			codes.Internal,
			"Invalid ObjectID converted",
		)
	}
	data := &BlogItem{}
	filter := bson.M{"_id" : objectId}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil{
		return nil, status.Errorf(
			codes.NotFound,
			"Couldnt find blog with id provided",
		)
	}

	return documentToBlog(data),nil

}
