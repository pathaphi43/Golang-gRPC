package main

import (
	"context"
	"log"
	"net"
	"net/http"
	pb "piefiredire/proto"
	pie "piefiredire/server/src"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedPieFireDireServiceServer
}

func (s *server) GetSummary(ctx context.Context, request *pb.GetSummaryRequest) (*pb.GetSummaryResponse, error) {
	summary := pie.GetBeefSummary()
	return &pb.GetSummaryResponse{Beef: summary}, nil
}

func main() {
	// flag.Parse()
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterPieFireDireServiceServer(grpcServer, &server{})
	// Serve gRPC server
	go func() {
		log.Printf("Server gRPC listening at %v", lis.Addr())
		log.Fatal(grpcServer.Serve(lis))
	}()

	mux := runtime.NewServeMux()
	err = pb.RegisterPieFireDireServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("Failed to register gRPC-Gateway: %v", err)
	}

	// Serve HTTP server
	log.Println("Serving HTTP on localhost:8080")
	http.ListenAndServe("localhost:8080", mux)
}
