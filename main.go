package main

import (
	"log"
	"net"

	grpcAPI "github.com/renantarouco/ics-name-server/api/grpc"
	"github.com/renantarouco/ics-name-server/api/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterNameServiceServer(grpcServer, &grpcAPI.NameServiceServer{})
	lis, err := net.Listen("tcp", "localhost:7001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)
}
