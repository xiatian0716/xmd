/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"{{.gomodinit}}/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var serviceHost = "127.0.0.1:8001"
	r := gin.Default()
	r.GET("/ping", callRPC)
	r.Run(serviceHost) //默认是8080端口
}

func callRPC(c *gin.Context) {
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials()) //不使用任何加密
	conn, err := grpc.Dial("127.0.0.1:8002", credentials)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	client := proto.New{{.projectname}}Client(conn)
	r, err := client.Ping(context.Background(), &proto.Request{Ping: "bobby"})
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"message": r.Pong,
	})
	fmt.Println(r.Pong)
}
