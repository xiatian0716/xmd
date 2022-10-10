/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package app

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/xiatian0716/xmd/goweb/util/app/dto"
)

func GetParams(c *gin.Context) dto.BasePage {
	var (
		page   int
		size   int
		blurry string
	)

	page = com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	size = com.StrTo(c.DefaultQuery("size", "1")).MustInt()
	blurry = c.DefaultQuery("blurry", "")

	return dto.BasePage{Page: page, Size: size, Blurry: blurry}
}
