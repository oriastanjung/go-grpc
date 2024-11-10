package main

import (
	"context"
	"fmt"
	pb "grpc_tutorial/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) GetAllBlogs(in *emptypb.Empty, stream grpc.ServerStreamingServer[pb.Blog]) error{
	log.Printf("get all blogs invoked")

	cursor, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil{
		return status.Errorf(codes.Internal,fmt.Sprintf("Error while find collection : %v",err))
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()){
		data := &BlogItem{}
		err := cursor.Decode(data)

		if err!=nil{
			return status.Errorf(codes.Internal,fmt.Sprintf("Error while decode data : %v",err))
		}
		stream.Send(documentToBlog(data))
	}

	if err := cursor.Err(); err!=nil{
		return status.Errorf(codes.Internal,fmt.Sprintf("Error while decoding from mongodb : %v",err))
	}
	return nil
}