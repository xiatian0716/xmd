/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package service

import (
	"context"
	"fmt"
	"{{.gomodinit}}/proto"
	"{{.gomodinit}}/rpc/app/models"
)

type Server struct {
	proto.Unimplemented{{.projectname}}Server
}

func (s *Server) Ping(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	fmt.Println("request:", req.Ping)
	pongs := models.GetPing()
	fmt.Println("response:", pongs.Name)
	return &proto.Response{Pong: "Hello " + pongs.Name}, nil
}
