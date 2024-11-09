package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Define the metadata key and value
const (
    requiredMetadataKey   = "authorization"
    requiredMetadataValue = "test-token"
)
// Unary interceptor for checking metadata
func UnaryMetadataInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    if err := checkMetadata(ctx); err != nil {
        return nil, err
    }
    return handler(ctx, req)
}

// Stream interceptor for checking metadata
func StreamMetadataInterceptor(
    srv interface{},
    ss grpc.ServerStream,
    info *grpc.StreamServerInfo,
    handler grpc.StreamHandler,
) error {
    if err := checkMetadata(ss.Context()); err != nil {
        return err
    }
    return handler(srv, ss)
}

// Helper function to check metadata
func checkMetadata(ctx context.Context) error {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return status.Errorf(codes.Unauthenticated, "missing metadata")
    }

    if values := md[requiredMetadataKey]; len(values) == 0 || values[0] != requiredMetadataValue {
        return status.Errorf(codes.Unauthenticated, "invalid or missing token")
    }

    return nil
}
