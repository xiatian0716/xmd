/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xiatian0716/xmd/goweb/app/service"
	"github.com/xiatian0716/xmd/goweb/util/app"
	"net/http"
)

// 部门api
type PingController struct {
}

func (e *PingController) GetPings(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.DefaultQuery("name", "")
	pingService := service.Ping{Name: name}
	vo := pingService.GetPings()
	appG.Response(http.StatusOK, app.SUCCESS, vo)
}

//func (e *PingController) Post(c *gin.Context) {
//	var (
//		//param params.XXXParam
//		model models.SysPing
//		appG  = app.Gin{C: c}
//	)
//	httpCode, errCode := app.BindAndValid(c, &model)
//	if errCode != app.SUCCESS {
//		appG.Response(httpCode, errCode, nil)
//		return
//	}
//	pingService := service.Ping{
//		Name: model.Name,
//	}
//
//	if err := pingService.PostPings(); err != nil {
//		appG.Response(http.StatusInternalServerError, app.FAIL_ADD_DATA, nil)
//		return
//	}
//
//	appG.Response(http.StatusOK, app.SUCCESS, nil)
//}
