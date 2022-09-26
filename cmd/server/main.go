package main

import (
	"fmt"
	"log"
	"net"
	"os"

	geomServer "github.com/by-sabbir/grpc-service-example/internal/geometry"
	pb "github.com/by-sabbir/grpc-service-example/proto"
	"google.golang.org/grpc"
)

var (
	host = "localhost"
	port = "5000"
)

func main() {
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}
	log.Println("tcp listener started at port: ", port)
	grpcServer := grpc.NewServer()
	geomServiceServer := geomServer.NewServer()
	// registering geometry service server into grpc server
	pb.RegisterGeometryServiceServer(grpcServer, geomServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}
