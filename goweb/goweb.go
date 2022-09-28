/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package main

import (
	"context"
	"fmt"
	"github.com/xiatian0716/xmd/goweb/api"
	"github.com/xiatian0716/xmd/goweb/util/snowflakex"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	//"github.com/xiatian0716/xmd/goweb/logic"
	"github.com/xiatian0716/xmd/goweb/util/conf"
	"github.com/xiatian0716/xmd/goweb/util/mysqlx"
	"github.com/xiatian0716/xmd/goweb/util/redisxx"
	"github.com/xiatian0716/xmd/goweb/util/zapx"
)

// GoWeb开发较通用的脚手架模板
func main() {
	// 1. 加载配置
	if err := conf.Setup("."); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	fmt.Println(conf.Conf)
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

	// 5. 注册路由
	r := api.Setup(conf.Conf.App.Mode)
	// 6. 启动服务（优雅关机）
	fmt.Println(conf.Conf.App.Port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Conf.App.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
