// cmd/server/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	db "github.com/ganduumar/klausapp-softwareengineer-test-task/internal/db"
	pb "github.com/ganduumar/klausapp-softwareengineer-test-task/internal/grpc/pb"
	s "github.com/ganduumar/klausapp-softwareengineer-test-task/internal/grpc/services"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	db, err := db.NewDB()
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRatingServiceServer(grpcServer, s.NewRatingsService(db))
	grpcServer.Serve(lis)
}
