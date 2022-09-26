package geometry

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/by-sabbir/grpc-service-example/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	srv := NewServer()
	pb.RegisterGeometryServiceServer(s, srv)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
func TestArea(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewGeometryServiceClient(conn)
	in := &pb.RectRequest{Height: 100.0, Width: 20.0}
	resp, err := client.Area(ctx, in)
	assert.NoError(t, err)
	area := in.Height * in.Width
	assert.Equal(t, area, resp.Result)
}

func TestPerimeter(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewGeometryServiceClient(conn)
	in := &pb.RectRequest{Height: 100.0, Width: 20.0}
	resp, err := client.Perimeter(ctx, in)
	assert.NoError(t, err)
	perim := 2 * (in.Height + in.Width)
	assert.Equal(t, perim, resp.Result)
}
