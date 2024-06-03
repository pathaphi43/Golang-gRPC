package main

import (
	"context"
	"flag"
	"log"

	pb "piefiredire/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewPieFireDireServiceClient(conn)

	res, err := client.GetSummary(context.Background(), &pb.GetSummaryRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(res.GetBeef())
}
