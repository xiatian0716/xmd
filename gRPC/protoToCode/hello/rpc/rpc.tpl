package main

import (
	"context"
	"fmt"
	"net"
	"{{.gomodinit}}/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.Unimplemented{{.projectname}}Server
}

func (s *server) Ping(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	fmt.Println("request:", req.Ping)
	return &proto.Response{Pong: "Hello " + req.Ping}, nil
}

func main() {
	g := grpc.NewServer()
	proto.Register{{.projectname}}Server(g, &server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8002")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
