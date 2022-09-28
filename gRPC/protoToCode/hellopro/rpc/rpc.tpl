/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package main

import (
	"fmt"
	"{{.gomodinit}}/proto"
	"{{.gomodinit}}/rpc/app/service"
	"github.com/xiatian0716/xmd/goweb/util/conf"
	"github.com/xiatian0716/xmd/goweb/util/mysqlx"
	"github.com/xiatian0716/xmd/goweb/util/redisxx"
	"github.com/xiatian0716/xmd/goweb/util/snowflakex"
	"github.com/xiatian0716/xmd/goweb/util/zapx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

// GoWeb开发较通用的脚手架模板
func main() {
	// 1. 加载配置
	if err := conf.Setup("./rpc"); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	// 2. 初始化日志
	if err := zapx.Setup(conf.Conf.Log, conf.Conf.App.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	// 3. 初始化MySQL连接
	if err := mysqlx.Setup(conf.Conf.MySQL); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysqlx.MysqlClose()
	// 4. 初始化Redis连接
	if err := redisxx.Setup(conf.Conf.Redis); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redisxx.RedisClose()

	if err := snowflakex.Setup(conf.Conf.Snowflake.StartTime, conf.Conf.Snowflake.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 5. 注册rpc
	g := grpc.NewServer()
	proto.Register{{.projectname}}Server(g, &service.Server{})

	// 6. 启动服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Conf.App.Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	defer g.Stop()
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

	zap.L().Info("Server exiting")
}
