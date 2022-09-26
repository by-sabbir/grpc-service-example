package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/by-sabbir/grpc-service-example/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GeometryServiceServer
}

func NewServer() *Server {
	return &Server{}
}

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
	geomServiceServer := NewServer()
	// registering gemoetry service server into grpc server
	pb.RegisterGeometryServiceServer(grpcServer, geomServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}

func (s *Server) Area(ctx context.Context, in *pb.RectRequest) (*pb.AreaResponse, error) {
	log.Println("invoked Area: ", in)
	return &pb.AreaResponse{
		Result: in.Height * in.Width,
	}, nil
}

func (s *Server) Perimeter(ctx context.Context, in *pb.RectRequest) (*pb.PermiterResponse, error) {
	log.Println("invoked Area: ", in)
	return &pb.PermiterResponse{
		Result: 2 * (in.Height + in.Width),
	}, nil
}
