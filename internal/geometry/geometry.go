package geometry

import (
	"context"
	"log"

	pb "github.com/by-sabbir/grpc-service-example/proto"
)

type Store interface {
	Area(context.Context, *pb.RectRequest) (*pb.AreaResponse, error)
	Perimeter(context.Context, *pb.RectRequest) (*pb.PermiterResponse, error)
}

type Server struct {
	pb.GeometryServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Area(ctx context.Context, in *pb.RectRequest) (*pb.AreaResponse, error) {
	log.Println("invoked Area: ", in)
	return &pb.AreaResponse{
		Result: in.Height * in.Width,
	}, nil
}

func (s *Server) Perimeter(ctx context.Context, in *pb.RectRequest) (*pb.PermiterResponse, error) {
	log.Println("invoked Perimeter: ", in)
	return &pb.PermiterResponse{
		Result: 2 * (in.Height + in.Width),
	}, nil
}
