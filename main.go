package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	shops "shops/proto"
)

type Good struct {
	shops.UnimplementedGoodServer
}

func (c *Good) Hellow(ctx context.Context, res *shops.HeRequest) (*shops.HeResponse, error) {
	return &shops.HeResponse{Name: "ffff"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Println(err.Error())
	}

	server := grpc.NewServer()
	shops.RegisterGoodServer(server, &Good{})
	reflection.Register(server)
	server.Serve(lis)

}
