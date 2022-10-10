/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xiatian0716/xmd/goweb/app/controllers"
	"github.com/xiatian0716/xmd/goweb/util/conf"
	"github.com/xiatian0716/xmd/goweb/util/zapx"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(zapx.GinLogger(), zapx.GinRecovery(true))

	// WebPing
	r.GET("/version", func(c *gin.Context) {
		// zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", "err"))
		c.String(http.StatusOK, conf.Conf.App.Version)
	})
	// 1.加载前端
	// 注意：路径是从/目录加载，而不是api目录
	// r.LoadHTMLFiles("./view/FormDataIndex.html")
	r.LoadHTMLGlob("view/*") // 更好用
	// r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Ping.html", nil)
	})

	// 2.api接口
	v1 := r.Group("/v1")

	// PingController
	// http://ip:8084/v1/ping
	pingController := controllers.PingController{}
	v1.GET("/ping", pingController.GetPings)

	return r
}
