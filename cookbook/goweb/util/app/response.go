/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponsePage struct {
	Code      int         `json:"status"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	TotalPage int         `json:"totalPage"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, errCode interface{}, data interface{}) {
	switch errCode.(type) {
	case int:
		intCode := errCode.(int)
		g.C.JSON(httpCode, Response{
			Code: intCode,
			Msg:  GetMsg(intCode),
			Data: data,
		})
	case string:
		strCode := errCode.(string)
		g.C.JSON(httpCode, Response{
			Code: 9999,
			Msg:  strCode,
			Data: data,
		})
	}

	return
}

// Response setting gin.JSON
func (g *Gin) ResponsePage(httpCode int, errCode interface{}, data interface{}, total, totalPage int) {
	intCode := errCode.(int)
	g.C.JSON(httpCode, ResponsePage{
		Code:      intCode,
		Msg:       GetMsg(intCode),
		Data:      data,
		Total:     total,
		TotalPage: totalPage,
	})
	return
}
